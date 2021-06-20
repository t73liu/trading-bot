package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/t73liu/tradingbot/lib/alpaca"
	"github.com/t73liu/tradingbot/lib/finviz"
	"github.com/t73liu/tradingbot/lib/newsapi"
	"github.com/t73liu/tradingbot/lib/options"
	"github.com/t73liu/tradingbot/lib/utils"
	"github.com/t73liu/tradingbot/lib/yahoo-finance"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type config struct {
	port         int
	dbURL        string
	sessionKey   string
	newsApiKey   string
	alpacaConfig alpaca.ClientConfig
}

type trader struct {
	conf          *config
	handler       *http.Handler
	sessionStore  *sessions.CookieStore
	logger        *zap.SugaredLogger
	db            *pgxpool.Pool
	newsClient    *newsapi.Client
	alpacaClient  *alpaca.Client
	yahooClient   *yahoo.Client
	finvizClient  *finviz.Client
	optionsClient *options.Client
}

const userIDContextKey = utils.ContextKey("userID")

var (
	invalidCredentialsError = errors.New("invalid credentials")
	unauthenticatedError    = errors.New("unauthenticated")
)

func main() {
	var conf config
	flag.IntVar(&conf.port, "port", 8080, "Server port")
	flag.StringVar(&conf.dbURL, "db.url", "", "URL to connect to traderdb")
	flag.StringVar(
		&conf.sessionKey,
		"session.key",
		"s6v9y$B&E)H+MbQeThWmZq4t7w!z%C*F",
		"32 byte authentication key used for cookies",
	)
	flag.StringVar(&conf.newsApiKey, "news.key", "", "News API Key")
	flag.StringVar(
		&conf.alpacaConfig.ApiKey,
		"alpaca.key",
		"",
		"Alpaca API Key",
	)
	flag.StringVar(
		&conf.alpacaConfig.ApiSecret,
		"alpaca.secret",
		"",
		"Alpaca API Secret",
	)
	flag.BoolVar(
		&conf.alpacaConfig.IsLiveTrading,
		"alpaca.live-trading",
		false,
		"Use Alpaca for live trading",
	)
	flag.BoolVar(
		&conf.alpacaConfig.IsPaidData,
		"alpaca.paid-data",
		false,
		"Use Alpaca with paid market data plan",
	)
	flag.Parse()

	t := newTrader(&conf)
	defer t.db.Close()

	t.logger.Info("Starting service with HTTP at port %s\n", conf.port)
	server := utils.NewHttpServer(conf.port, t.handler)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			t.logger.Fatal(err)
		}
	}()

	// Gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownContext); err != nil {
		t.logger.Fatal("shutdown error", err)
	}
	t.logger.Info("Shutting down")
}

func newTrader(conf *config) *trader {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	zLogger, err := zapConfig.Build()
	if err != nil {
		log.Fatalln("Failed to initialize logger:", err)
	}

	t := &trader{
		conf:   conf,
		logger: zLogger.Sugar(),
	}

	t.logger.Info("Initializing ...")
	if conf.dbURL == "" {
		t.logger.Fatal("-db.url flag must be provided")
	}
	if conf.newsApiKey == "" {
		t.logger.Fatal("-news.key flag must be provided")
	}
	if conf.alpacaConfig.ApiKey == "" {
		t.logger.Fatal("-alpaca.key flag must be provided")
	}
	if conf.alpacaConfig.ApiSecret == "" {
		t.logger.Fatal("-alpaca.secret flag must be provided")
	}
	if len(conf.sessionKey) < 32 {
		t.logger.Fatal("-session.key flag must be provided and 32 bytes long")
	}

	t.db, err = pgxpool.Connect(context.Background(), conf.dbURL)
	if err != nil {
		t.logger.Fatal("Unable to connect to database:", err)
	}

	t.sessionStore = sessions.NewCookieStore([]byte(conf.sessionKey))
	t.sessionStore.MaxAge(int(12 * time.Hour / time.Second))

	// Initialize 3rd party API clients
	client := utils.NewHttpClient()
	t.newsClient = newsapi.NewClient(client, t.conf.newsApiKey)
	t.conf.alpacaConfig.HttpClient = client
	t.alpacaClient = alpaca.NewClient(t.conf.alpacaConfig)
	t.yahooClient = yahoo.NewClient(client)
	t.finvizClient = finviz.NewClient(client)
	t.optionsClient = options.NewClient(client)

	// Setup HTTP router
	router := mux.NewRouter()
	router.Use(
		utils.PanicRecovery,
		t.logRequests,
		utils.SecureHeaders,
	)
	router.PathPrefix("/assets/").Handler(http.StripPrefix(
		"/assets/",
		http.FileServer(http.Dir("assets/")),
	))
	t.addAuthRoutes(router)

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(t.requireAuthentication)
	t.addNewsRoutes(apiRouter.PathPrefix("/news").Subrouter())
	t.addStockRoutes(apiRouter.PathPrefix("/stocks").Subrouter())
	t.addAccountRoutes(apiRouter.PathPrefix("/account").Subrouter())

	router.PathPrefix("/").Handler(http.HandlerFunc(t.spaHandler))
	return t
}

func (t *trader) spaHandler(w http.ResponseWriter, r *http.Request) {
	if _, authenticated := t.getSessionUserID(r); !authenticated {
		if r.URL.Path != "/login" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
	}
	http.ServeFile(w, r, "assets/index.html")
}

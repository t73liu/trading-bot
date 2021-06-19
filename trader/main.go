package main

import (
	"context"
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
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type config struct {
	port         int
	dbURL        string
	newsApiKey   string
	alpacaConfig alpaca.ClientConfig
}

type trader struct {
	conf          *config
	handler       *http.Handler
	logger        *zap.SugaredLogger
	db            *pgxpool.Pool
	newsClient    *newsapi.Client
	alpacaClient  *alpaca.Client
	yahooClient   *yahoo.Client
	finvizClient  *finviz.Client
	optionsClient *options.Client
}

// TODO Remove hardcoded userID
const userID = 1

func main() {
	var conf config
	flag.IntVar(&conf.port, "port", 8080, "Server port")
	flag.StringVar(&conf.dbURL, "db.url", "", "URL to connect to traderdb")
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

	t.db, err = pgxpool.Connect(context.Background(), conf.dbURL)
	if err != nil {
		t.logger.Fatal("Unable to connect to database:", err)
	}

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
	router.PathPrefix("/assets/").Handler(http.StripPrefix(
		"/assets/",
		http.FileServer(http.Dir("assets/")),
	))
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/index.html")
	})
	t.AddNewsRoutes(router.PathPrefix("/api/news").Subrouter())
	t.AddStockRoutes(router.PathPrefix("/api/stocks").Subrouter())
	t.AddAccountRoutes(router.PathPrefix("/api/account").Subrouter())

	// Setup middleware
	router.Use(LogResponseTime(t.logger))

	return t
}

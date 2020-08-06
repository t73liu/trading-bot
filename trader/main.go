package main

import (
	"context"
	"flag"
	"github.com/caddyserver/certmagic"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"strings"
	"tradingbot/lib/newsapi"
	"tradingbot/lib/polygon"
	"tradingbot/lib/utils"
	"tradingbot/lib/yahoo-finance"
	"tradingbot/trader/account"
	"tradingbot/trader/news"
	"tradingbot/trader/stock"
)

func main() {
	prodFlag := flag.Bool("prod", false, "Run in production mode")
	httpsFlag := flag.Bool("https", false, "Run with HTTPS")
	emailFlag := flag.String(
		"email",
		"",
		"Email to receive expiration alerts for certificates (Optional)",
	)
	domainsFlag := flag.String(
		"domains",
		"",
		"Comma-delimited domains (Required for HTTPS)",
	)
	flag.Parse()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	logger.Println("Initializing ...")

	client := utils.NewHttpClient()

	dbPool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatalln("Unable to connect to database:", err)
	}
	defer dbPool.Close()

	handler := initApp(logger, client, dbPool)

	if *httpsFlag {
		// https://github.com/caddyserver/certmagic#requirements
		logger.Println("Starting service with HTTPS")
		domainNames := strings.Split(strings.TrimSpace(*domainsFlag), ",")
		if len(domainNames) == 0 {
			logger.Fatalln("domains are required for HTTPS")
		}
		certmagic.DefaultACME.Agreed = true
		email := strings.TrimSpace(*emailFlag)
		if email != "" {
			certmagic.DefaultACME.Email = email
		}
		if !*prodFlag {
			certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA
		}
		logger.Fatalln(certmagic.HTTPS(domainNames, handler))
	}

	port := ":8080"
	logger.Printf("Starting service with HTTP at port %s\n", port)
	server := utils.NewHttpServer(port, &handler)
	logger.Fatalln(server.ListenAndServe())
}

func initApp(logger *log.Logger, client *http.Client, db *pgxpool.Pool) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/index.html")
	})
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	newsClient := newsapi.NewClient(client, os.Getenv("NEWS_API_KEY"))
	newsHandlers := news.NewHandlers(logger, db, newsClient)
	newsHandlers.AddRoutes(router)

	polygonClient := polygon.NewClient(client, os.Getenv("ALPACA_API_KEY"))
	yahooClient := yahoo.NewClient(client)
	stockHandlers := stock.NewHandlers(logger, db, polygonClient, yahooClient)
	stockHandlers.AddRoutes(router)

	accountHandlers := account.NewHandlers(logger, db)
	accountHandlers.AddRoutes(router)

	return router
}

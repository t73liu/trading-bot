package main

import (
	"context"
	"flag"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"tradingbot/lib/newsapi"
	"tradingbot/lib/polygon"
	"tradingbot/lib/utils"
	"tradingbot/lib/yahoo-finance"
	"tradingbot/trader/account"
	"tradingbot/trader/middleware"
	"tradingbot/trader/news"
	"tradingbot/trader/stocks"
)

func main() {
	portFlag := flag.String("port", ":8080", "Server port")
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

	port := *portFlag
	logger.Printf("Starting service with HTTP at port %s\n", port)
	server := utils.NewHttpServer(port, &handler)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatalln(err)
		}
	}()

	// Gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownContext); err != nil {
		logger.Fatalln("shutdown error", err)
	}
	logger.Println("Shutting down ...")
}

func initApp(logger *log.Logger, client *http.Client, db *pgxpool.Pool) http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/index.html")
	})

	newsClient := newsapi.NewClient(client, os.Getenv("NEWS_API_KEY"))
	newsHandlers := news.NewHandlers(logger, db, newsClient)
	newsHandlers.AddRoutes(router.PathPrefix("/api/news").Subrouter())

	polygonClient := polygon.NewClient(client, os.Getenv("ALPACA_API_KEY"))
	yahooClient := yahoo.NewClient(client)
	stockHandlers := stocks.NewHandlers(logger, db, polygonClient, yahooClient)
	stockHandlers.AddRoutes(router.PathPrefix("/api/stocks").Subrouter())

	accountHandlers := account.NewHandlers(logger, db)
	accountHandlers.AddRoutes(router.PathPrefix("/api/account").Subrouter())

	router.Use(middleware.LogResponseTime(logger))

	return router
}

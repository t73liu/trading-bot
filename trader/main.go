package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"github.com/mholt/certmagic"
	"github.com/t73liu/trading-bot/trader/news"
	"github.com/t73liu/trading-bot/trader/stock"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	logger.Println("Initializing ...")

	// prod := flag.Bool("prod", false, "Run in production mode")
	domains := flag.String("domains", "", "Use HTTPS for the following domain(s)")
	flag.Parse()

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets/index.html")
	})
	router.ServeFiles("/assets/*filepath", http.Dir("assets/"))

	client := http.Client{Timeout: 15 * time.Second}

	newsClient := news.NewClient(client, os.Getenv("NEWS_API_KEY"))
	newsHandlers := news.NewHandlers(logger, newsClient)
	newsHandlers.AddRoutes(router)

	stockHandlers := stock.NewHandlers(logger)
	stockHandlers.AddRoutes(router)

	if *domains != "" {
		// https://github.com/mholt/certmagic#requirements
		logger.Printf("Starting service with HTTPS\n")
		logger.Fatal(certmagic.HTTPS(strings.Split(*domains, ","), router))
	}
	logger.Printf("Starting service with HTTP at port %s\n", ":80")
	logger.Fatal(http.ListenAndServe(":80", router))
}

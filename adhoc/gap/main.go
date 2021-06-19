package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/t73liu/tradingbot/lib/finviz"
	"github.com/t73liu/tradingbot/lib/newsapi"
	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dbURL := flag.String("db.url", "", "URL to connect to traderdb")
	newsApiKey := flag.String("news.key", "", "News API Key")
	alpacaApiKey := flag.String("alpaca.key", "", "Alpaca API Key")
	flag.Parse()

	if *dbURL == "" {
		log.Fatalln("-db.url flag must be provided")
	}
	if *newsApiKey == "" {
		log.Fatalln("-news.key flag must be provided")
	}
	if *alpacaApiKey == "" {
		log.Fatalln("-alpaca.key flag must be provided")
	}

	db, err := pgxpool.Connect(context.Background(), *dbURL)
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	tradableSymbols, err := traderdb.GetStocksBySymbol(db)
	if err != nil {
		log.Fatalln("Failed to query tradable stocks from DB:", err)
	}

	httpClient := utils.NewHttpClient()
	finvizClient := finviz.NewClient(httpClient)
	gapStocks, err := finvizClient.ScreenStocksOverview("v=111&f=ta_gap_u7&ft=4&o=-gap")
	if err != nil {
		log.Fatalln("Failed to screen for gap stocks:", err)
	}
	tradableGapStocks := make([]finviz.StockOverview, 0, len(gapStocks))
	for _, gapStock := range gapStocks {
		if _, ok := tradableSymbols[gapStock.Symbol]; ok {
			tradableGapStocks = append(tradableGapStocks, gapStock)
			log.Printf("%+v\n", gapStock)
		}
	}

	now := time.Now()
	newsClient := newsapi.NewClient(httpClient, *newsApiKey)

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalln("Failed to load location America/New_York:", err)
	}

	startTime := utils.GetLastWeekday(now)

	for _, stock := range tradableGapStocks {
		response, err := newsClient.GetAllHeadlinesWithSources(newsapi.AllArticlesQueryParams{
			Query:     utils.TrimCompanyName(stock.Company) + " OR " + stock.Symbol + " Stock",
			StartTime: startTime,
			Domains:   utils.NewsDomains,
			PageSize:  10,
		})
		if err != nil {
			log.Fatalln("Failed to fetch news for stocks:", err)
		}
		for _, article := range response.Articles {
			log.Println(article.Title)
			log.Println("URL:", article.Url)
			log.Println("Published At:", article.PublishedAt.In(location))
		}
	}
}

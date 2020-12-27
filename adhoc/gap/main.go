package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"tradingbot/lib/finviz"
	"tradingbot/lib/newsapi"
	"tradingbot/lib/polygon"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	databaseUrl := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL environment variable is required")
		os.Exit(1)
	}
	newsAPIKey := strings.TrimSpace(os.Getenv("NEWS_API_KEY"))
	if newsAPIKey == "" {
		fmt.Println("NEWS_API_KEY environment variable is required")
		os.Exit(1)
	}
	alpacaAPIKey := strings.TrimSpace(os.Getenv("ALPACA_API_KEY"))
	if alpacaAPIKey == "" {
		fmt.Println("ALPACA_API_KEY environment variable is required")
		os.Exit(1)
	}

	db, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	tradableSymbols, err := traderdb.GetStocksBySymbol(db)
	if err != nil {
		fmt.Println("Failed to query tradable stocks from DB:", err)
		os.Exit(1)
	}

	httpClient := utils.NewHttpClient()
	polygonClient := polygon.NewClient(httpClient, alpacaAPIKey)
	movers, err := polygonClient.GetMovers(true)
	if err != nil {
		fmt.Println("Failed to fetch movers:", err)
		os.Exit(1)
	}
	tradableMovers := make([]traderdb.Stock, 0, len(movers))
	for _, mover := range movers {
		if stock, ok := tradableSymbols[mover.Ticker]; ok {
			tradableMovers = append(tradableMovers, stock)
			fmt.Printf("%+v\n", mover)
		}
	}

	finvizClient := finviz.NewClient(httpClient)
	gapStocks, err := finvizClient.ScreenStocksOverview("v=111&f=ta_gap_u7&ft=4&o=-gap")
	if err != nil {
		fmt.Println("Failed to screen for gap stocks:", err)
		os.Exit(1)
	}
	tradableGapStocks := make([]finviz.StockOverview, 0, len(gapStocks))
	for _, gapStock := range gapStocks {
		if _, ok := tradableSymbols[gapStock.Symbol]; ok {
			tradableGapStocks = append(tradableGapStocks, gapStock)
			fmt.Printf("%+v\n", gapStock)
		}
	}

	now := time.Now()
	newsClient := newsapi.NewClient(httpClient, newsAPIKey)

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Failed to load location America/New_York:", err)
		os.Exit(1)
	}

	startTime := utils.GetLastWeekday(now)

	for _, stock := range tradableMovers {
		response, err := newsClient.GetAllHeadlinesWithSources(newsapi.AllArticlesQueryParams{
			Query:     utils.TrimCompanyName(stock.Company) + " OR " + stock.Symbol + " Stock",
			StartTime: startTime,
			Domains:   utils.NewsDomains,
			PageSize:  10,
		})
		if err != nil {
			fmt.Println("Failed to fetch news for stocks:", err)
			os.Exit(1)
		}
		for _, article := range response.Articles {
			fmt.Println(article.Title)
			fmt.Println("URL:", article.Url)
			fmt.Println("Published At:", article.PublishedAt.In(location))
		}
	}
}

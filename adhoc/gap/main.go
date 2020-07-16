package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/t73liu/trading-bot/lib/finviz"
	"github.com/t73liu/trading-bot/lib/newsapi"
	"github.com/t73liu/trading-bot/lib/polygon"
	"github.com/t73liu/trading-bot/lib/traderdb"
	"github.com/t73liu/trading-bot/lib/utils"
	"os"
	"strings"
	"time"
)

// TODO Dedupe
var domains = []string{
	"finance.yahoo.com",
	"investors.com",
	"seekingalpha.com",
	"fool.com",
	"reuters.com",
	"bloomberg.com",
}

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

	stocks, err := traderdb.GetTradableStocks(db)
	if err != nil {
		fmt.Println("Failed to query tradable stocks from DB:", err)
		os.Exit(1)
	}
	tradableSymbols := make(map[string]traderdb.Stock)
	for _, stock := range stocks {
		tradableSymbols[stock.Symbol] = stock
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
	gapStocks, err := finvizClient.ScreenStocks("v=111&f=ta_gap_u7&ft=4&o=-gap")
	if err != nil {
		fmt.Println("Failed to screen for gap stocks:", err)
		os.Exit(1)
	}
	tradableGapStocks := make([]finviz.StockInfo, 0, len(gapStocks))
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

	//for _, stock := range tradableMovers {
	for _, stock := range tradableGapStocks {
		response, err := newsClient.GetAllHeadlinesBySources(newsapi.AllArticlesQueryParams{
			Query:     trimCompanyName(stock.Company) + " OR " + stock.Symbol + " Stock",
			StartTime: startTime,
			Domains:   domains,
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

// TODO Dedupe
func trimCompanyName(company string) string {
	trimmedName := strings.TrimSpace(strings.Split(company, " Class ")[0])
	//trimmedName = strings.TrimSpace(strings.Split(company, " Inc.")[0])
	return trimmedName
}

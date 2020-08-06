package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
	"strings"
	"time"
	"tradingbot/adhoc/test-strategy/strategy"
	"tradingbot/lib/candle"
	"tradingbot/lib/traderdb"
)

func main() {
	databaseUrl := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL environment variable is required")
		os.Exit(1)
	}

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Failed to loading America/New_York timezone:", err)
		os.Exit(1)
	}

	startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, location)
	endTime := time.Date(2020, 8, 1, 0, 0, 0, 0, location)
	candles, err := traderdb.GetStockCandles(conn, "SHOP", startTime, endTime)
	if err != nil {
		fmt.Println("Failed to fetch stock candles from DB:", err)
		os.Exit(1)
	}

	formattedCandles, err := candle.CompressCandles(
		candle.FillMinuteCandles(candle.FilterTradingHourCandles(candles)),
		5,
		"minute",
		location,
	)
	if err != nil {
		fmt.Println("Failed to format stock candles:", err)
		os.Exit(1)
	}
	applyStrategy(formattedCandles, 10000)
}

func applyStrategy(candles []candle.Candle, capital float64) {
	capitalMicros := candle.DollarsToMicros(capital)
	//dailyPortfolios := strategy.Hold(candles, capitalMicros)
	//dailyPortfolios := strategy.TrailingStop(candles, capitalMicros, 0.95)
	dailyPortfolios := strategy.RSI(candles, capitalMicros, 70, 30)
	for _, portfolio := range dailyPortfolios {
		strategy.PrintPortfolio(portfolio)
	}
}

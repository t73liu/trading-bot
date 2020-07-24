package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	analyze "github.com/t73liu/trading-bot/lib/technical-analysis"
	"github.com/t73liu/trading-bot/lib/traderdb"
	"os"
	"strings"
	"time"
	"trading-bot/adhoc/test-strategy/strategy"
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

	candles, err := traderdb.GetTradingHourStockCandles(conn, "SHOP")
	if err != nil {
		fmt.Println("Failed to fetch stock candles from DB:", err)
		os.Exit(1)
	}

	formattedCandles, err := analyze.CompressCandles(
		analyze.FillMinuteCandles(convertCandles(candles)),
		1,
		"minute",
		location,
	)
	if err != nil {
		fmt.Println("Failed to format stock candles:", err)
		os.Exit(1)
	}
	applyStrategy(formattedCandles, 10000)
}

// TODO move to shared module
func convertCandles(candles []traderdb.Candle) []analyze.Candle {
	results := make([]analyze.Candle, 0, len(candles))
	for _, candle := range candles {
		results = append(results, analyze.Candle{
			OpenedAt: candle.OpenedAt,
			Volume:   candle.Volume,
			Open:     candle.OpenMicros,
			High:     candle.HighMicros,
			Low:      candle.LowMicros,
			Close:    candle.CloseMicros,
		})
	}
	return results
}

func applyStrategy(candles []analyze.Candle, capital float64) {
	capitalMicros := analyze.DollarsToMicros(capital)
	//dailyPortfolios := strategy.Hold(candles, capitalMicros)
	dailyPortfolios := strategy.TrailingStop(candles, capitalMicros, 0.95)
	for _, portfolio := range dailyPortfolios {
		strategy.PrintPortfolio(portfolio)
	}
}

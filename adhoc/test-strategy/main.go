package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/t73liu/tradingbot/adhoc/test-strategy/strategy"
	"github.com/t73liu/tradingbot/lib/candle"
	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/jackc/pgx/v4"
)

func main() {
	dbURL := flag.String("db.url", "", "URL to connect to traderdb")
	flag.Parse()

	if *dbURL == "" {
		log.Fatalln("-db.url flag must be provided")
	}

	conn, err := pgx.Connect(context.Background(), *dbURL)
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalln("Failed to loading America/New_York timezone:", err)
	}

	startTime := time.Date(2020, 1, 1, 0, 0, 0, 0, location)
	endTime := time.Date(2020, 8, 1, 0, 0, 0, 0, location)
	candles, err := traderdb.GetStockCandles(conn, "SHOP", startTime, endTime)
	if err != nil {
		log.Fatalln("Failed to fetch stock candles from DB:", err)
	}

	formattedCandles, err := candle.CompressCandles(
		candle.FillMinuteCandles(candle.FilterTradingHourCandles(candles)),
		5,
		"minute",
		location,
	)
	if err != nil {
		log.Fatalln("Failed to format stock candles:", err)
	}
	applyStrategy(formattedCandles, 10000)
}

func applyStrategy(candles []candle.Candle, capital float64) {
	capitalMicros := utils.DollarsToMicros(capital)
	//dailyPortfolios := strategy.Hold(candles, capitalMicros)
	//dailyPortfolios := strategy.TrailingStop(candles, capitalMicros, 0.95)
	dailyPortfolios := strategy.RSI(candles, capitalMicros, 70, 30)
	for _, portfolio := range dailyPortfolios {
		strategy.PrintPortfolio(portfolio)
	}
}

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

	applyStrategy(
		analyze.CompressCandles(
			analyze.FillMinuteCandles(convertCandles(candles)),
			1,
			"minute",
			location,
		),
	)
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

func applyStrategy(candles []analyze.Candle) {
	if len(candles) == 0 {
		return
	}

	dates := make([]string, 0)
	candlesByDay := make(map[string][]analyze.Candle)
	for _, candle := range candles {
		date := candle.OpenedAt.Format("2006-01-02")
		groupedCandles, ok := candlesByDay[date]
		if !ok {
			dates = append(dates, date)
		}
		candlesByDay[date] = append(groupedCandles, candle)
	}

	capital := int64(0)
	buyPrice := candles[0].Close
	numberOfShares := int64(100)
	hasLongPosition := true
	initialCapital := analyze.MicrosToDollars(buyPrice * numberOfShares)
	trades := 0
	fmt.Println("INITIAL CAPITAL", initialCapital)
	for _, date := range dates {
		groupedCandles := candlesByDay[date]
		fmt.Println("Trades on day:", date, len(groupedCandles))
		closingPrices := getClosingPrices(groupedCandles)
		//macds := analyze.StandardMACD(closingPrices)
		fasts := analyze.EMA(closingPrices, 20)
		//slows := analyze.SMA(closingPrices, 50)
		//rsiValues := analyze.RSI(closingPrices, 14)
		for i, candle := range groupedCandles {
			//currentRSI := rsiValues[i]
			fast := fasts[i]
			if fast.Valid && candle.Close < fast.Micro {
				if hasLongPosition {
					fmt.Printf(
						"Potential sell on %s at %f with RSI %f\n",
						candle.OpenedAt.Format("15:04:05"),
						analyze.MicrosToDollars(candle.Close),
						fast.Micro,
					)
					capital += candle.Close * numberOfShares
					numberOfShares = 0
					hasLongPosition = false
					fmt.Println(analyze.MicrosToDollars(capital))
					trades++
				}
			}
			if fast.Valid && candle.Close > fast.Micro {
				if !hasLongPosition {
					fmt.Printf(
						"BUY: %s at %.2f with RSI %.2f\n",
						candle.OpenedAt.Format("15:04:05"),
						analyze.MicrosToDollars(candle.Close),
						fast.Micro,
					)
					hasLongPosition = true
					buyPrice = candle.Close
					numberOfShares = capital / buyPrice
					capital -= numberOfShares * buyPrice
					fmt.Println(numberOfShares, "BOUGHT", analyze.MicrosToDollars(buyPrice))
				}
			}
		}
	}

	fmt.Println(
		"Hold",
		analyze.MicrosToDollars(candles[len(candles)-1].Close*100),
		analyze.MicrosToDollars(candles[len(candles)-1].Close*100)/initialCapital*100,
	)
	fmt.Println(
		"END CAPITAL",
		analyze.MicrosToDollars(buyPrice*numberOfShares+capital),
		analyze.MicrosToDollars(buyPrice*numberOfShares+capital)/initialCapital*100,
		trades,
	)
}

func getClosingPrices(candles []analyze.Candle) (results []int64) {
	for _, candle := range candles {
		results = append(results, candle.Close)
	}
	return results
}

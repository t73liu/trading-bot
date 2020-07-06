package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	analyze "github.com/t73liu/trading-bot/lib/technical-analysis"
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

	candles := getStockCandles(conn, "SHOP")
	applyStrategy(compressCandles(fillMissingCandles(candles), 1))
}

const stockCandlesQuery = `
SELECT opened_at, open_micros, high_micros, low_micros, close_micros, volume FROM stock_candles
WHERE stock_id = $1
ORDER BY opened_at
`

func getStockCandles(conn *pgx.Conn, symbol string) (candles []analyze.Candle) {
	var stockId int
	err := conn.QueryRow(
		context.Background(),
		"SELECT id FROM stocks WHERE symbol = $1",
		symbol,
	).Scan(&stockId)
	if err != nil {
		fmt.Println("Failed to fetch stock with symbol:", symbol, err)
		os.Exit(1)
	}

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Failed to load time location:", err)
		os.Exit(1)
	}
	rows, err := conn.Query(context.Background(), stockCandlesQuery, stockId)
	if err != nil {
		fmt.Println("Failed to fetch stock candles:", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var openMicros int64
		var highMicros int64
		var lowMicros int64
		var closeMicros int64
		var volume int64
		var openedAt time.Time
		if err = rows.Scan(
			&openedAt, &openMicros, &highMicros, &lowMicros, &closeMicros, &volume,
		); err != nil {
			fmt.Println("Failed to parse row:", err)
			os.Exit(1)
		}
		openedAt = openedAt.In(location)
		if isWithinTradingHours(openedAt) {
			candles = append(candles, analyze.Candle{
				OpenedAt: openedAt,
				Open:     openMicros,
				High:     highMicros,
				Low:      lowMicros,
				Close:    closeMicros,
				Volume:   volume,
			})
		}
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Failed to parse rows:", err)
		os.Exit(1)
	}

	return candles
}

func isWithinTradingHours(moment time.Time) bool {
	hour, minute, _ := moment.Clock()
	if hour == 9 {
		return minute >= 30
	}
	return hour > 9 && hour < 16
}

func fillMissingCandles(candles []analyze.Candle) []analyze.Candle {
	filledCandles := make([]analyze.Candle, 0, len(candles))
	var prevCandle analyze.Candle
	var prevTime time.Time
	for i, candle := range candles {
		currentTime := candle.OpenedAt
		if i > 0 && prevTime.Day() == currentTime.Day() {
			minutesDiff := int(currentTime.Sub(prevTime).Minutes())
			for i := minutesDiff; i > 1; i-- {
				backfilledTime := currentTime.Add(-1 * time.Minute * time.Duration(i-1))
				filledCandles = append(filledCandles, genPlaceholderCandle(prevCandle, backfilledTime))
			}
		}
		prevCandle = candle
		prevTime = prevCandle.OpenedAt
		filledCandles = append(filledCandles, candle)
	}
	return filledCandles
}

func genPlaceholderCandle(candle analyze.Candle, openedAt time.Time) analyze.Candle {
	return analyze.Candle{
		OpenedAt: openedAt,
		Volume:   0,
		Open:     candle.Close,
		High:     candle.Close,
		Low:      candle.Close,
		Close:    candle.Close,
	}
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

// TODO need to respect day separation and non-zero starts
func compressCandles(candles []analyze.Candle, interval int) []analyze.Candle {
	if interval == 1 {
		return candles
	}
	newCandles := make([]analyze.Candle, 0, len(candles)/interval)
	count := 0
	for _, candle := range candles {
		if count == interval-1 {
			newCandles = append(newCandles, candle)
			count = 0
		} else {
			count++
		}
	}
	return newCandles
}

func getClosingPrices(candles []analyze.Candle) (results []int64) {
	for _, candle := range candles {
		results = append(results, candle.Close)
	}
	return results
}

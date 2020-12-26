package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"tradingbot/lib/alpaca"
	"tradingbot/lib/traderdb"
	"tradingbot/lib/utils"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

const userId = 1

func main() {
	databaseUrl := strings.TrimSpace(os.Getenv("DATABASE_URL"))
	if databaseUrl == "" {
		fmt.Println("DATABASE_URL environment variable is required")
		os.Exit(1)
	}
	apiKey := strings.TrimSpace(os.Getenv("ALPACA_API_KEY"))
	if apiKey == "" {
		fmt.Println("ALPACA_API_KEY environment variable is required")
		os.Exit(1)
	}
	apiSecretKey := strings.TrimSpace(os.Getenv("ALPACA_API_SECRET"))
	if apiSecretKey == "" {
		fmt.Println("ALPACA_API_SECRET environment variable is required")
		os.Exit(1)
	}

	db, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		os.Exit(1)
	}

	alpacaClient := alpaca.NewClient(
		utils.NewHttpClient(),
		apiKey,
		apiSecretKey,
		false,
	)

	stocks, err := traderdb.GetWatchlistStocksByUserId(db, userId)
	if err != nil {
		fmt.Println("Failed to fetch watchlist stocks:", err)
		os.Exit(1)
	}

	stockSymbols := make([]string, 0, len(stocks))
	for _, stock := range stocks {
		stockSymbols = append(stockSymbols, stock.Symbol)
	}
	now := time.Now()
	candlesBySymbol, err := alpacaClient.GetCandles(alpaca.CandleQueryParams{
		Symbols:    stockSymbols,
		CandleSize: alpaca.OneMin,
		StartTime:  now.AddDate(-1, 0, 0),
		EndTime:    now,
	})

	if err = bulkInsertStockCandles(db, candlesBySymbol, stocks); err != nil {
		fmt.Println("Failed to bulk insert stock candles:", err)
		os.Exit(1)
	}
}

func bulkInsertStockCandles(
	db *pgxpool.Pool,
	candlesBySymbol map[string][]alpaca.Candle,
	stocks []traderdb.Stock,
) error {
	tx, err := db.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	rows := make([][]interface{}, 0, len(stocks))
	for _, stock := range stocks {
		if candles, ok := candlesBySymbol[stock.Symbol]; ok {
			for _, candle := range candles {
				rows = append(rows, []interface{}{
					stock.Id,
					utils.ConvertUnixSecondsToTime(candle.StartAtUnixSec),
					convertFloatToMicros(candle.Open),
					convertFloatToMicros(candle.High),
					convertFloatToMicros(candle.Low),
					convertFloatToMicros(candle.Close),
					candle.Volume,
				})
			}
		}
	}
	_, err = tx.CopyFrom(
		context.Background(),
		pgx.Identifier{"stock_candles"},
		[]string{
			"stock_id",
			"opened_at",
			"open_micros",
			"high_micros",
			"low_micros",
			"close_micros",
			"volume",
		},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

func convertFloatToMicros(number float32) int64 {
	return int64(number * 1000000)
}

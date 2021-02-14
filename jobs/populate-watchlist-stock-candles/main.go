package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
	"tradingbot/lib/alpaca"
	"tradingbot/lib/candle"
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

	stocks, err := traderdb.GetAllWatchlistStocks(db)
	if err != nil {
		fmt.Println("Failed to fetch watchlist stocks:", err)
		os.Exit(1)
	}

	stockSymbols := make([]string, 0, len(stocks))
	for _, stock := range stocks {
		stockSymbols = append(stockSymbols, stock.Symbol)
	}
	now := time.Now()
	candlesBySymbol, err := alpacaClient.GetCandlesBySymbol(alpaca.CandleQueryParams{
		Symbols:    stockSymbols,
		Limit:      1000,
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

	var candles []candle.Candle
	for _, stock := range stocks {
		for _, c := range candlesBySymbol[stock.Symbol] {
			candles = append(candles, candle.Candle{
				StockID:     int64(stock.ID),
				OpenedAt:    utils.ConvertUnixSecondsToTime(c.StartAtUnixSec),
				Volume:      int64(c.Volume),
				OpenMicros:  utils.DollarsToMicros(float64(c.Open)),
				HighMicros:  utils.DollarsToMicros(float64(c.High)),
				LowMicros:   utils.DollarsToMicros(float64(c.Low)),
				CloseMicros: utils.DollarsToMicros(float64(c.Close)),
			})
		}
	}

	if err = traderdb.UpsertStockCandles(tx, &candles); err != nil {
		return err
	}

	if err = tx.Commit(context.Background()); err != nil {
		return err
	}

	return nil
}

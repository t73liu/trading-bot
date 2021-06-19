package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/t73liu/tradingbot/lib/alpaca"
	"github.com/t73liu/tradingbot/lib/candle"
	"github.com/t73liu/tradingbot/lib/traderdb"
	"github.com/t73liu/tradingbot/lib/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	dbURL := flag.String("db.url", "", "URL to connect to traderdb")
	alpacaApiKey := flag.String("alpaca.key", "", "Alpaca API Key")
	alpacaApiSecret := flag.String("alpaca.secret", "", "Alpaca API Secret")
	flag.Parse()

	if *dbURL == "" {
		log.Fatalln("-db.url flag must be provided")
	}
	if *alpacaApiKey == "" {
		log.Fatalln("-alpaca.key flag must be provided")
	}
	if *alpacaApiSecret == "" {
		log.Fatalln("-alpaca.secret flag must be provided")
	}

	db, err := pgxpool.Connect(context.Background(), *dbURL)
	if err != nil {
		log.Fatalln("Failed to connect to DB:", err)
	}

	alpacaClient := alpaca.NewClient(alpaca.ClientConfig{
		HttpClient:    utils.NewHttpClient(),
		ApiKey:        *alpacaApiKey,
		ApiSecret:     *alpacaApiSecret,
		IsLiveTrading: false,
		IsPaidData:    false,
	})

	stocks, err := traderdb.GetAllWatchlistStocks(db)
	if err != nil {
		log.Fatalln("Failed to fetch watchlist stocks:", err)
	}

	candlesBySymbol := make(map[string][]alpaca.Candle)
	for _, stock := range stocks {
		now := time.Now()
		candlesResponse, err := alpacaClient.GetSymbolCandles(stock.Symbol, alpaca.CandleQueryParams{
			Limit:      10000,
			CandleSize: alpaca.OneMin,
			StartTime:  now.AddDate(-1, 0, 0),
			EndTime:    now,
		})
		if err != nil {
			log.Fatalln(fmt.Sprintf("Failed to fetch %s candles:", stock.Symbol), err)
		}
		candlesBySymbol[stock.Symbol] = candlesResponse.Candles
	}

	if err = bulkInsertStockCandles(db, candlesBySymbol, stocks); err != nil {
		log.Fatalln("Failed to bulk insert stock candles:", err)
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
				OpenedAt:    c.StartAt,
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

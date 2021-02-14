package traderdb

import (
	"context"
	"fmt"
	"strings"
	"time"
	"tradingbot/lib/candle"
	"tradingbot/lib/utils"
)

const stockCandlesQuery = `
SELECT opened_at, open_micros, high_micros, low_micros, close_micros, volume FROM stock_candles
WHERE stock_id = $1 AND opened_at BETWEEN $2 AND $3
ORDER BY opened_at
`

func GetStockCandles(db PGConnection, symbol string, startTime time.Time, endTime time.Time) (candles []candle.Candle, err error) {
	var stockID int64
	err = db.QueryRow(
		context.Background(),
		"SELECT id FROM stocks WHERE symbol = $1",
		symbol,
	).Scan(&stockID)
	if err != nil {
		return candles, err
	}

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return candles, err
	}

	rows, err := db.Query(context.Background(), stockCandlesQuery, stockID, startTime, endTime)
	if err != nil {
		return candles, err
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
			return candles, err
		}
		openedAt = openedAt.In(location)
		candles = append(candles, candle.Candle{
			StockID:     stockID,
			OpenedAt:    openedAt,
			OpenMicros:  openMicros,
			HighMicros:  highMicros,
			LowMicros:   lowMicros,
			CloseMicros: closeMicros,
			Volume:      volume,
		})
	}

	if err = rows.Err(); err != nil {
		return candles, err
	}

	return candles, nil
}

const upsertStockCandlesQuery = `
INSERT INTO stock_candles
(
  stock_id,
  opened_at,
  open_micros,
  high_micros,
  low_micros,
  close_micros,
  volume
)
VALUES %s
ON CONFLICT (stock_id, opened_at)
DO UPDATE SET
  open_micros = EXCLUDED.open_micros,
  high_micros = EXCLUDED.high_micros,
  low_micros = EXCLUDED.low_micros,
  close_micros = EXCLUDED.close_micros,
  volume = EXCLUDED.volume
`

// Batched because PostgreSQL has a  limit of 65000 parameters
func UpsertStockCandles(db PGConnection, candles *[]candle.Candle) error {
	if candles == nil || len(*candles) == 0 {
		return nil
	}
	numberOfParams := 7
	batchedCandles := batchCandles(candles, 5000)
	for _, batch := range *batchedCandles {
		var values []interface{}
		var insertStr string
		for i, c := range batch {
			values = append(
				values,
				c.StockID,
				c.OpenedAt,
				c.OpenMicros,
				c.HighMicros,
				c.LowMicros,
				c.CloseMicros,
				c.Volume,
			)
			insertStr += utils.CreatePositionalParams(i*numberOfParams+1, numberOfParams)
		}
		query := fmt.Sprintf(upsertStockCandlesQuery, strings.TrimSuffix(insertStr, ","))
		if _, err := db.Exec(context.Background(), query, values...); err != nil {
			return err
		}
	}
	return nil
}

func batchCandles(candles *[]candle.Candle, batchSize int) *[][]candle.Candle {
	batches := len(*candles) / batchSize
	if len(*candles)%batchSize != 0 {
		batches++
	}
	batchedCandles := make([][]candle.Candle, batches)

	for i, c := range *candles {
		batchedCandles[i/batchSize] = append(batchedCandles[i/batchSize], c)
	}
	return &batchedCandles
}

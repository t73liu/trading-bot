package traderdb

import (
	"context"
	"time"
	"tradingbot/lib/candle"
)

const stockCandlesQuery = `
SELECT opened_at, open_micros, high_micros, low_micros, close_micros, volume FROM stock_candles
WHERE stock_id = $1 AND opened_at BETWEEN $2 AND $3
ORDER BY opened_at
`

func GetStockCandles(db PGConnection, symbol string, startTime time.Time, endTime time.Time) (candles []candle.Candle, err error) {
	var stockID int
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

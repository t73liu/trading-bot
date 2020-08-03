package traderdb

import (
	"context"
	"time"
)

type Candle struct {
	OpenedAt    time.Time
	OpenMicros  int64
	HighMicros  int64
	LowMicros   int64
	CloseMicros int64
	Volume      int64
}

const stockCandlesQuery = `
SELECT opened_at, open_micros, high_micros, low_micros, close_micros, volume FROM stock_candles
WHERE stock_id = $1
ORDER BY opened_at
`

func GetTradingHourStockCandles(db PGConnection, symbol string) (candles []Candle, err error) {
	var stockId int
	err = db.QueryRow(
		context.Background(),
		"SELECT id FROM stocks WHERE symbol = $1",
		symbol,
	).Scan(&stockId)
	if err != nil {
		return candles, err
	}

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return candles, err
	}

	rows, err := db.Query(context.Background(), stockCandlesQuery, stockId)
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
		if IsWithinNATradingHours(openedAt) {
			candles = append(candles, Candle{
				OpenedAt:    openedAt,
				OpenMicros:  openMicros,
				HighMicros:  highMicros,
				LowMicros:   lowMicros,
				CloseMicros: closeMicros,
				Volume:      volume,
			})
		}
	}

	if err = rows.Err(); err != nil {
		return candles, err
	}

	return candles, nil
}

// Assuming location = "America/New_York"
func IsWithinNATradingHours(moment time.Time) bool {
	hour, minute, _ := moment.Clock()
	if hour == 9 {
		return minute >= 30
	}
	return hour > 9 && hour < 16
}

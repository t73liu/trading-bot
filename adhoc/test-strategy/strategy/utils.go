package strategy

import analyze "github.com/t73liu/trading-bot/lib/technical-analysis"

func groupCandlesByDate(candles []analyze.Candle) ([]string, map[string][]analyze.Candle) {
	dates := make([]string, 0)
	candlesByDate := make(map[string][]analyze.Candle)
	for _, candle := range candles {
		date := candle.OpenedAt.Format("2006-01-02")
		groupedCandles, ok := candlesByDate[date]
		if !ok {
			dates = append(dates, date)
		}
		candlesByDate[date] = append(groupedCandles, candle)
	}

	return dates, candlesByDate
}

func calcPercentChange(prevMicros int64, currentMicros int64) float64 {
	ratio := analyze.MicrosToDollars(currentMicros) / analyze.MicrosToDollars(prevMicros)
	return analyze.RoundToTwoDecimals(ratio*100 - 100)
}

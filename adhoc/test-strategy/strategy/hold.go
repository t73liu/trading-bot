package strategy

import (
	analyze "github.com/t73liu/trading-bot/lib/technical-analysis"
)

func Hold(candles []analyze.Candle, capitalMicros int64) []Portfolio {
	if len(candles) < 2 {
		return nil
	}
	firstCandle := candles[0]
	shares := capitalMicros / firstCandle.Open
	dates, candlesByDate := groupCandlesByDate(candles)
	dailySnapshots := make([]Portfolio, 0, len(dates)+1)
	initialSnapshot := Portfolio{
		Date:               "initial",
		Cash:               capitalMicros - firstCandle.Open*shares,
		SharesHeld:         shares,
		EndOfDayValue:      capitalMicros,
		DailyChange:        0,
		DailyPercentChange: 0,
		AllTimePerformance: 0,
	}
	prevSnapshot := initialSnapshot
	dailySnapshots = append(dailySnapshots, initialSnapshot)
	for _, date := range dates {
		dailyCandles := candlesByDate[date]
		closingPrice := dailyCandles[len(dailyCandles)-1].Close
		marketValue := prevSnapshot.Cash + shares*closingPrice
		snapshot := Portfolio{
			Date:               date,
			Cash:               prevSnapshot.Cash,
			SharePrice:         closingPrice,
			SharesHeld:         shares,
			EndOfDayValue:      marketValue,
			DailyChange:        marketValue - prevSnapshot.EndOfDayValue,
			DailyPercentChange: calcPercentChange(prevSnapshot.EndOfDayValue, marketValue),
			AllTimePerformance: calcPercentChange(capitalMicros, marketValue),
		}
		dailySnapshots = append(dailySnapshots, snapshot)
		prevSnapshot = snapshot
	}
	return dailySnapshots
}

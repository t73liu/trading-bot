package strategy

import (
	"github.com/t73liu/tradingbot/lib/candle"
)

func Hold(candles []candle.Candle, capitalMicros int64) []Portfolio {
	if len(candles) < 2 {
		return nil
	}
	firstCandle := candles[0]
	dates, candlesByDate := groupCandlesByDate(candles)
	dailySnapshots := make([]Portfolio, 0, len(dates)+1)
	prevSnapshot := genInitialPortfolio(capitalMicros, firstCandle.OpenMicros)
	dailySnapshots = append(dailySnapshots, prevSnapshot)
	for _, date := range dates {
		dailyCandles := candlesByDate[date]
		closingPrice := dailyCandles[len(dailyCandles)-1].CloseMicros
		marketValue := prevSnapshot.Cash + prevSnapshot.SharesHeld*closingPrice
		snapshot := Portfolio{
			Date:               date,
			Cash:               prevSnapshot.Cash,
			SharePrice:         closingPrice,
			SharesHeld:         prevSnapshot.SharesHeld,
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

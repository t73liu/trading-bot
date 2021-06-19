package strategy

import (
	"github.com/t73liu/tradingbot/lib/candle"
)

// Sell when trailing loss percent is reached and buy back at beginning of the next day
func TrailingStop(candles []candle.Candle, capitalMicros int64, stopLoss float64) []Portfolio {
	if len(candles) < 2 {
		return nil
	}
	firstCandle := candles[0]
	dates, candlesByDate := groupCandlesByDate(candles)
	dailySnapshots := make([]Portfolio, 0, len(dates)+1)
	prevSnapshot := genInitialPortfolio(capitalMicros, firstCandle.OpenMicros)
	dailySnapshots = append(dailySnapshots, prevSnapshot)
	currentHigh := prevSnapshot.SharePrice
	for _, date := range dates {
		dailyCandles := candlesByDate[date]
		trades := make([]Trade, 0, 2)
		shares := prevSnapshot.SharesHeld
		currentCash := prevSnapshot.Cash
		// Buy at open if not currently holding any shares
		if prevSnapshot.SharesHeld == 0 {
			firstCandle = dailyCandles[0]
			openPrice := firstCandle.OpenMicros
			shares = currentCash / openPrice
			currentCash -= openPrice * shares
			currentHigh = openPrice
			trades = append(trades, Trade{
				Type:           Buy,
				NumberOfShares: shares,
				PriceMicros:    openPrice,
				Timestamp:      firstCandle.OpenedAt,
			})
		}

		// Check trailing stop loss percent
		for _, dailyCandle := range dailyCandles {
			if currentHigh < dailyCandle.HighMicros {
				currentHigh = dailyCandle.HighMicros
			}
			sharePrice := dailyCandle.CloseMicros
			trailingStopLoss := float64(sharePrice) / float64(currentHigh)
			if trailingStopLoss <= stopLoss {
				trades = append(trades, Trade{
					Type:           Sell,
					NumberOfShares: shares,
					PriceMicros:    sharePrice,
					Timestamp:      dailyCandle.OpenedAt,
				})
				// Approximation of sale price
				currentCash += shares * sharePrice
				shares = 0
				break
			}
		}

		closingPrice := dailyCandles[len(dailyCandles)-1].CloseMicros
		if shares == 0 {
			closingPrice = 0
		}
		marketValue := currentCash + shares*closingPrice
		snapshot := Portfolio{
			Date:               date,
			Cash:               currentCash,
			SharePrice:         closingPrice,
			SharesHeld:         shares,
			EndOfDayValue:      marketValue,
			DailyChange:        marketValue - prevSnapshot.EndOfDayValue,
			DailyPercentChange: calcPercentChange(prevSnapshot.EndOfDayValue, marketValue),
			AllTimePerformance: calcPercentChange(capitalMicros, marketValue),
			Trades:             trades,
		}
		dailySnapshots = append(dailySnapshots, snapshot)
		prevSnapshot = snapshot
	}
	return dailySnapshots
}

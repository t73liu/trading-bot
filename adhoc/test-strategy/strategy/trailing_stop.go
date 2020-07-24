package strategy

import (
	analyze "github.com/t73liu/trading-bot/lib/technical-analysis"
)

// Sell when trailing loss percent is reached and buy back at beginning of the next day
func TrailingStop(candles []analyze.Candle, capitalMicros int64, stopLoss float64) []Portfolio {
	if len(candles) < 2 {
		return nil
	}
	firstCandle := candles[0]
	dates, candlesByDate := groupCandlesByDate(candles)
	dailySnapshots := make([]Portfolio, 0, len(dates)+1)
	prevSnapshot := genInitialPortfolio(capitalMicros, firstCandle.Open)
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
			shares = prevSnapshot.Cash / firstCandle.Open
			currentCash -= firstCandle.Open * shares
			currentHigh = firstCandle.Open
			trades = append(trades, Trade{
				Type:           Buy,
				NumberOfShares: shares,
				PriceMicros:    firstCandle.Open,
				Timestamp:      firstCandle.OpenedAt,
			})
		}

		// Check trailing stop loss percent
		for _, candle := range dailyCandles {
			if currentHigh < candle.High {
				currentHigh = candle.High
			}
			trailingStopLoss := float64(candle.Close) / float64(currentHigh)
			if trailingStopLoss <= stopLoss {
				trades = append(trades, Trade{
					Type:           Sell,
					NumberOfShares: shares,
					PriceMicros:    candle.Close,
					Timestamp:      candle.OpenedAt,
				})
				// Approximation of sale price
				currentCash += shares * candle.Close
				shares = 0
				break
			}
		}

		closingPrice := dailyCandles[len(dailyCandles)-1].Close
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

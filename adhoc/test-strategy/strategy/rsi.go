package strategy

import (
	"fmt"
	"tradingbot/lib/candle"
	analyze "tradingbot/lib/technical-analysis"
)

// Buy when RSI crosses above lower limit and sell when RSI crosses below upper limit
func RSI(candles []candle.Candle, capitalMicros int64, upperLimit, lowerLimit float64) []Portfolio {
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
		trades := make([]Trade, 0, 2)
		shares := prevSnapshot.SharesHeld
		currentCash := prevSnapshot.Cash

		// Check for RSI crossovers
		rsiValues := analyze.RSI(candle.GetClosingPrices(dailyCandles), 14)
		for i, dailyCandle := range dailyCandles {
			if prevIndex := i - 1; prevIndex > 0 && rsiValues[prevIndex].Valid && rsiValues[i].Valid {
				prevRSIValue := rsiValues[prevIndex].Value
				currRSIValue := rsiValues[i].Value
				sharePrice := dailyCandle.CloseMicros
				if shares == 0 {
					// Buy if RSI increases above lower limit (i.e. oversold)
					if prevRSIValue <= lowerLimit && currRSIValue > lowerLimit {
						shares = currentCash / sharePrice
						currentCash -= sharePrice * shares
						trades = append(trades, Trade{
							Type:           Buy,
							NumberOfShares: shares,
							PriceMicros:    sharePrice,
							Timestamp:      dailyCandle.OpenedAt,
							Details:        fmt.Sprintf("RSI %.2f to %.2f", prevRSIValue, currRSIValue),
						})
					}
				} else {
					// Sell if RSI decreases below upper limit (i.e. overbought)
					if prevRSIValue >= upperLimit && currRSIValue < upperLimit {
						trades = append(trades, Trade{
							Type:           Sell,
							NumberOfShares: shares,
							PriceMicros:    sharePrice,
							Timestamp:      dailyCandle.OpenedAt,
							Details:        fmt.Sprintf("RSI %.2f to %.2f", prevRSIValue, currRSIValue),
						})
						currentCash += sharePrice * shares
						shares = 0
					}
				}
			}
		}

		closingPrice := dailyCandles[len(dailyCandles)-1].CloseMicros
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

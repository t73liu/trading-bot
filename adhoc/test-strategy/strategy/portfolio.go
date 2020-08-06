package strategy

import (
	"fmt"
	"tradingbot/lib/candle"
)

type Portfolio struct {
	Date               string
	Cash               int64
	SharePrice         int64
	SharesHeld         int64
	Trades             []Trade
	EndOfDayValue      int64
	DailyChange        int64
	DailyPercentChange float64
	AllTimePerformance float64
}

func PrintPortfolio(portfolio Portfolio) {
	fmt.Printf(
		"Date: %s, Cash: %.2f, Shares: %d, Share Price: %.2f, Value: %.2f,"+
			" Daily Change: %.2f, Daily Percent Change: %.2f%%, All Time: %.2f%%\n",
		portfolio.Date,
		candle.MicrosToDollars(portfolio.Cash),
		portfolio.SharesHeld,
		candle.MicrosToDollars(portfolio.SharePrice),
		candle.MicrosToDollars(portfolio.EndOfDayValue),
		candle.MicrosToDollars(portfolio.DailyChange),
		portfolio.DailyPercentChange,
		portfolio.AllTimePerformance,
	)
	for _, trade := range portfolio.Trades {
		fmt.Println("  Trade:", trade)
	}
}

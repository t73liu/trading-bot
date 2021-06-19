package strategy

import (
	"log"

	"github.com/t73liu/tradingbot/lib/utils"
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
	log.Printf(
		"Date: %s, Cash: %.2f, Shares: %d, Share Price: %.2f, Value: %.2f,"+
			" Daily Change: %.2f, Daily Percent Change: %.2f%%, All Time: %.2f%%\n",
		portfolio.Date,
		utils.MicrosToDollars(portfolio.Cash),
		portfolio.SharesHeld,
		utils.MicrosToDollars(portfolio.SharePrice),
		utils.MicrosToDollars(portfolio.EndOfDayValue),
		utils.MicrosToDollars(portfolio.DailyChange),
		portfolio.DailyPercentChange,
		portfolio.AllTimePerformance,
	)
	for _, trade := range portfolio.Trades {
		log.Println("  Trade:", trade)
	}
}

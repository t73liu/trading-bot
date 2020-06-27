package main

import (
	"fmt"
	"github.com/t73liu/trading-bot/lib/yahoo-stock-calendar"
	"net/http"
	"os"
	"time"
)

// TODO send email prior to market hours with news for watchlist stocks and
// upcoming IPOs and Earnings
func main() {
	client := yahoo.NewClient(&http.Client{Timeout: 15 * time.Second})
	now := time.Now()

	// Get Earnings/IPOs for the next 2 weeks
	for days := 0; days < 14; days++ {
		date := now.AddDate(0, 0, days)
		earnings, err := client.GetEarningsCall(date)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, earningsCall := range earnings {
			fmt.Printf("%+v\n", earningsCall)
		}

		ipos, err := client.GetIPOs(date)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, ipo := range ipos {
			fmt.Printf("%+v\n", ipo)
		}
	}
}

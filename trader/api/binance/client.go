package binance

import (
	"fmt"
	"github.com/t73liu/trading-bot/trader/candle"
)

func GetCandles(symbol string) []candle.Candle {
	var candles []candle.Candle
	fmt.Println(symbol)
	return candles
}

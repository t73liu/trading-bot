package crypto

import "github.com/t73liu/trading-bot/trader/api/binance"

func InsertCandles(symbol string) {
	binance.GetCandles(symbol)
}

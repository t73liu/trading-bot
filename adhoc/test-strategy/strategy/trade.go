package strategy

import (
	"fmt"
	analyze "github.com/t73liu/trading-bot/lib/technical-analysis"
	"time"
)

type TradeType string

const (
	Buy  TradeType = "Buy"
	Sell TradeType = "Sell"
)

type Trade struct {
	Type           TradeType
	NumberOfShares int64
	PriceMicros    int64
	Timestamp      time.Time
	Details        string
}

func (trade Trade) String() string {
	return fmt.Sprintf(
		"%s %d shares at %.2f - %s - %s",
		trade.Type,
		trade.NumberOfShares,
		analyze.MicrosToDollars(trade.PriceMicros),
		trade.Timestamp,
		trade.Details,
	)
}

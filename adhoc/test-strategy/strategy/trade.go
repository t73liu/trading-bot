package strategy

import "time"

type TradeType string

const (
	Sale TradeType = "Sale"
	Buy  TradeType = "Buy"
)

type Trade struct {
	Type           TradeType
	NumberOfShares int64
	PriceMicros    int64
	Timestamp      time.Time
}

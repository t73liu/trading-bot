package candle

import "time"

// OHLC values are in cents
type Candle struct {
	Open      float32   `json:"open"`
	High      float32   `json:"high"`
	Low       float32   `json:"low"`
	Close     float32   `json:"close"`
	Volume    uint32    `json:"volume"`
	StartTime time.Time `json:"startTime"`
}

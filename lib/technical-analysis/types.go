package analyze

import "time"

type Candle struct {
	OpenedAt time.Time
	Volume   int64
	Open     int64
	High     int64
	Low      int64
	Close    int64
}

type ValidMicro struct {
	Micro int64
	Valid bool
}

type ValidFloat struct {
	Value float64
	Valid bool
}

package alpaca

import "time"

type Asset struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Exchange     string `json:"exchange"`
	AssetClass   string `json:"class"`
	Symbol       string `json:"symbol"`
	Status       string `json:"status"`
	Tradable     bool   `json:"tradable"`
	Marginable   bool   `json:"marginable"`
	Shortable    bool   `json:"shortable"`
	EasyToBorrow bool   `json:"easy_to_borrow"`
	Fractionable bool   `json:"fractionable"`
}

type Candle struct {
	StartAt time.Time `json:"t"`
	Open    float32   `json:"o"`
	High    float32   `json:"h"`
	Low     float32   `json:"l"`
	Close   float32   `json:"c"`
	Volume  int32     `json:"v"`
}

type CandleSize string

const (
	OneMin  CandleSize = "1Min"
	OneHour CandleSize = "1Hour"
	OneDay  CandleSize = "1Day"
)

type LastTrade struct {
	Price             float32 `json:"price"`
	Size              int     `json:"size"`
	TimestampUnixNano int64   `json:"timestamp"`
}

type LastQuote struct {
	AskPrice          float32 `json:"askprice"`
	AskSize           int     `json:"asksize"`
	BidPrice          float32 `json:"bidprice"`
	BidSize           int     `json:"bidsize"`
	TimestampUnixNano int64   `json:"timestamp"`
}

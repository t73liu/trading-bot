package alpaca

type Asset struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Exchange     string `json:"exchange"`
	Class        string `json:"asset_class"`
	Symbol       string `json:"symbol"`
	Status       string `json:"status"`
	Tradable     bool   `json:"tradable"`
	Marginable   bool   `json:"marginable"`
	Shortable    bool   `json:"shortable"`
	EasyToBorrow bool   `json:"easy_to_borrow"`
}

type Candle struct {
	Time   int64   `json:"t"`
	Open   float32 `json:"o"`
	High   float32 `json:"h"`
	Low    float32 `json:"l"`
	Close  float32 `json:"c"`
	Volume int32   `json:"v"`
}

type CandleSize string

const (
	OneMin     CandleSize = "1Min"
	FiveMin    CandleSize = "5Min"
	FifteenMin CandleSize = "15Min"
	OneDay     CandleSize = "1D"
)

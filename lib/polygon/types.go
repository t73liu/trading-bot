package polygon

import "time"

type TickersResponse struct {
	Count   int      `json:"count"`
	Tickers []Ticker `json:"tickers"`
}

type Ticker struct {
	Ticker          string `json:"ticker"`
	Name            string `json:"name"`
	Market          string `json:"market"`
	Locale          string `json:"locale"`
	Currency        string `json:"currency"`
	Active          bool   `json:"active"`
	PrimaryExchange string `json:"primaryExch"`
}

type Article struct {
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Source    string    `json:"source"`
	Summary   string    `json:"summary"`
	Timestamp time.Time `json:"timestamp"`
}

type TickerDetails struct {
	Symbol      string   `json:"symbol"`
	Name        string   `json:"name"`
	Exchange    string   `json:"exchange"`
	Country     string   `json:"country"`
	Industry    string   `json:"industry"`
	Sector      string   `json:"sector"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	MarketCap   int64    `json:"marketcap"`
	Similar     []string `json:"similar"`
	Tags        []string `json:"tags"`
}

type TickerBarsResponse struct {
	Ticker       string      `json:"ticker"`
	Status       string      `json:"status"`
	Adjusted     bool        `json:"adjusted"`
	QueryCount   int         `json:"queryCount"`
	ResultsCount int         `json:"resultsCount"`
	Results      []TickerBar `json:"results"`
}

type TickerBar struct {
	Ticker            string  `json:"T"`
	Volume            float64 `json:"v"`
	Open              float64 `json:"o"`
	High              float64 `json:"h"`
	Low               float64 `json:"l"`
	Close             float64 `json:"c"`
	StartAtUnixMillis int64   `json:"t"`
}

type TickerOHLC struct {
	Volume float64 `json:"v"`
	Open   float32 `json:"o"`
	High   float32 `json:"h"`
	Low    float32 `json:"l"`
	Close  float32 `json:"c"`
}

type TickerSnapshot struct {
	Ticker             string     `json:"ticker"`
	Day                TickerOHLC `json:"day"`
	PrevDay            TickerOHLC `json:"prevDay"`
	Change             float32    `json:"todaysChange"`
	ChangePercent      float32    `json:"todaysChangePerc"`
	UpdatedAtUnixNanos int64      `json:"updated"`
}

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
	MarketCap   int      `json:"marketcap"`
	Similar     []string `json:"similar"`
	Tags        []string `json:"tags"`
}

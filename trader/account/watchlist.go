package account

type StockSymbol struct {
	Id     int    `json:"id"`
	Symbol string `json:"symbol"`
}

type Watchlist struct {
	Id     int           `json:"id"`
	Name   string        `json:"name"`
	Stocks []StockSymbol `json:"stocks"`
}

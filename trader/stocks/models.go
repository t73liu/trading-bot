package stocks

import "time"

type Detail struct {
	Price         float64     `json:"price"`
	Company       string      `json:"company"`
	Website       string      `json:"website"`
	Description   string      `json:"description"`
	Sector        string      `json:"sector"`
	Industry      string      `json:"industry"`
	Country       string      `json:"country"`
	AverageVolume int64       `json:"averageVolume"`
	MarketCap     int64       `json:"marketCap"`
	SimilarStocks []string    `json:"similarStocks"`
	Shortable     bool        `json:"shortable"`
	Marginable    bool        `json:"marginable"`
	News          interface{} `json:"news"`
}

type Snapshot struct {
	Symbol         string    `json:"symbol"`
	Company        string    `json:"company"`
	Change         float64   `json:"change"`
	ChangePercent  float64   `json:"changePercent"`
	PreviousVolume int64     `json:"previousVolume"`
	PreviousClose  float64   `json:"previousClose"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

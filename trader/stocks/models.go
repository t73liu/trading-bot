package stocks

import (
	"time"
	"tradingbot/lib/utils"
)

type Detail struct {
	Price         float64           `json:"price"`
	Company       string            `json:"company"`
	Website       *utils.NullString `json:"website"`
	Description   *utils.NullString `json:"description"`
	Sector        *utils.NullString `json:"sector"`
	Industry      *utils.NullString `json:"industry"`
	Country       *utils.NullString `json:"country"`
	AverageVolume int64             `json:"averageVolume"`
	MarketCap     int64             `json:"marketCap"`
	SimilarStocks []string          `json:"similarStocks"`
	Shortable     bool              `json:"shortable"`
	Marginable    bool              `json:"marginable"`
	News          interface{}       `json:"news"`
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

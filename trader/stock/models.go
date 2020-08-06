package stock

type Detail struct {
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

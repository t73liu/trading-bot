package finviz

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	client http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: *httpClient,
	}
}

type StockInfo struct {
	Symbol        string
	Company       string
	Sector        string
	Industry      string
	Country       string
	MarketCap     int64
	Price         float64
	PercentChange float64
	Volume        int64
}

func (c *Client) ScreenStocks(query string) (stocks []StockInfo, err error) {
	req, err := http.NewRequest("GET", "https://finviz.com/screener.ashx", nil)
	if err != nil {
		return stocks, err
	}
	req.URL.RawQuery = query

	resp, err := c.client.Do(req)
	if err != nil {
		return stocks, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return stocks, errors.New("Response failed with status code: " + resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return stocks, err
	}
	doc.Find("#screener-content tr table").Each(func(i int, table *goquery.Selection) {
		if i == 2 {
			var marketCap int64
			var price float64
			var percentChange float64
			var volume int
			table.Find("tr").EachWithBreak(func(i int, row *goquery.Selection) bool {
				if i != 0 {
					contents := row.Contents().Map(func(_ int, cell *goquery.Selection) string {
						return cell.Text()
					})
					marketCap, err = convertHumanizedString(contents[7])
					if err != nil {
						return false
					}
					price, err = strconv.ParseFloat(contents[9], 64)
					if err != nil {
						return false
					}
					percentChange, err = strconv.ParseFloat(strings.TrimSuffix(contents[10], "%"), 64)
					if err != nil {
						return false
					}
					volume, err = strconv.Atoi(strings.ReplaceAll(contents[11], ",", ""))
					if err != nil {
						return false
					}
					stocks = append(stocks, StockInfo{
						Symbol:        contents[2],
						Company:       contents[3],
						Sector:        contents[4],
						Industry:      contents[5],
						Country:       contents[6],
						MarketCap:     marketCap,
						Price:         price,
						PercentChange: percentChange,
						Volume:        int64(volume),
					})
				}
				return true
			})
		}
	})

	return stocks, err
}

func convertHumanizedString(str string) (int64, error) {
	if str == "-" {
		return 0, nil
	}
	if strings.Contains(str, "B") {
		billions, err := strconv.ParseFloat(strings.TrimSuffix(str, "B"), 64)
		if err != nil {
			return 0, err
		}
		return int64(billions * 1000000000), nil
	}
	millions, err := strconv.ParseFloat(strings.TrimSuffix(str, "M"), 64)
	if err != nil {
		return 0, err
	}
	return int64(millions * 1000000), nil
}

package yahoo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"tradingbot/lib/utils"
)

const baseURL = "https://finance.yahoo.com"

const maxEventsPerPage = 100

type EarningsCallTime string

const (
	TransferAgentSystem EarningsCallTime = "TAS"
	TimeNotSupplied     EarningsCallTime = "TNS"
	BeforeMarketOpen    EarningsCallTime = "BMO"
	AfterMarketClose    EarningsCallTime = "AFC"
)

type EarningsCall struct {
	Ticker             string
	Company            string
	StartTime          EarningsCallTime
	EPSEstimate        *float64
	EPSActual          *float64
	EPSSurprisePercent *float64
	QuoteType          string
	Date               string
}

type IPO struct {
	Ticker    string
	Company   string
	Exchange  string
	PriceFrom *float64
	PriceTo   *float64
	Currency  string
	QuoteType string
	Date      string
}

type Stock struct {
	Symbol        string
	Company       string
	Exchange      string
	MarketCap     int64
	Price         float64
	Sector        string
	Industry      string
	Description   string
	Country       string
	Website       string
	AverageVolume int64
	News          []Article
}

type Article struct {
	URL         string    `json:"url"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	PublishedAt time.Time `json:"publishedAt"`
}

type Client struct {
	client http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: *httpClient,
	}
}

func (c *Client) GetEarningsCall(date time.Time) (earnings []EarningsCall, err error) {
	formattedDate := formatISO(date)
	offset := 0
	for {
		rows, err := c.getEvents("earnings", date, offset)
		if err != nil {
			return earnings, err
		}
		for _, row := range rows {
			record := row.(map[string]interface{})
			earningsCall := EarningsCall{
				Ticker:             record["ticker"].(string),
				Company:            record["companyshortname"].(string),
				StartTime:          EarningsCallTime(record["startdatetimetype"].(string)),
				EPSEstimate:        getFloat64(record["epsestimate"]),
				EPSActual:          getFloat64(record["epsactual"]),
				EPSSurprisePercent: getFloat64(record["epssurprisepct"]),
				QuoteType:          record["quoteType"].(string),
				Date:               formattedDate,
			}
			earnings = append(earnings, earningsCall)
		}
		if len(rows) == maxEventsPerPage {
			offset += maxEventsPerPage
		} else {
			break
		}
	}
	return earnings, nil
}

func (c *Client) GetIPOs(date time.Time) (ipos []IPO, err error) {
	formattedDate := formatISO(date)
	offset := 0
	for {
		rows, err := c.getEvents("ipo", date, offset)
		if err != nil {
			return ipos, err
		}
		for _, row := range rows {
			record := row.(map[string]interface{})
			ipo := IPO{
				Ticker:    record["ticker"].(string),
				Company:   record["companyshortname"].(string),
				Exchange:  record["exchange_short_name"].(string),
				PriceFrom: getFloat64(record["pricefrom"]),
				PriceTo:   getFloat64(record["priceto"]),
				Currency:  record["currencyname"].(string),
				QuoteType: record["quoteType"].(string),
				Date:      formattedDate,
			}
			ipos = append(ipos, ipo)
		}
		if len(rows) == maxEventsPerPage {
			offset += maxEventsPerPage
		} else {
			break
		}
	}
	return ipos, err
}

func (c *Client) getEvents(eventType string, date time.Time, offset int) ([]interface{}, error) {
	req, err := http.NewRequest("GET", baseURL+"/calendar/"+eventType, nil)
	if err != nil {
		return nil, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("size", "100")
	if !date.IsZero() {
		queryParams.Add("day", formatISO(date))
	}
	if offset > 0 {
		queryParams.Add("offset", strconv.Itoa(offset))
	}
	req.URL.RawQuery = queryParams.Encode()

	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bodyString := string(bodyBytes)
	for _, str := range strings.Split(bodyString, "\n") {
		length := len(str)
		if length > 16 && str[:16] == "root.App.main = " {
			var data map[string]interface{}
			if err = json.Unmarshal([]byte(str[16:length-1]), &data); err != nil {
				return nil, err
			}
			context := data["context"].(map[string]interface{})
			dispatcher := context["dispatcher"].(map[string]interface{})
			stores := dispatcher["stores"].(map[string]interface{})
			resultsStore := stores["ScreenerResultsStore"].(map[string]interface{})
			results := resultsStore["results"].(map[string]interface{})
			rows := results["rows"].([]interface{})
			return rows, nil
		}
	}
	return nil, errors.New("could not parse Yahoo Finance response")
}

func (c *Client) GetStock(symbol string) (stock Stock, err error) {
	location := utils.GetNYSELocation()

	response, err := http.Get(baseURL + "/quote/" + symbol)
	if err != nil {
		return stock, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return stock, err
	}

	bodyString := string(bodyBytes)
	for _, str := range strings.Split(bodyString, "\n") {
		length := len(str)
		if length > 16 && str[:16] == "root.App.main = " {
			var data map[string]interface{}
			if err = json.Unmarshal([]byte(str[16:length-1]), &data); err != nil {
				return stock, err
			}
			context := data["context"].(map[string]interface{})
			dispatcher := context["dispatcher"].(map[string]interface{})
			stores := dispatcher["stores"].(map[string]interface{})
			quoteStore := stores["QuoteSummaryStore"].(map[string]interface{})
			// Price and volume details
			stockDetail := quoteStore["price"].(map[string]interface{})
			marketCap := stockDetail["marketCap"].(map[string]interface{})
			price := stockDetail["regularMarketPrice"].(map[string]interface{})
			volume := stockDetail["averageDailyVolume10Day"].(map[string]interface{})
			stock.Symbol = stockDetail["symbol"].(string)
			stock.Company = stockDetail["longName"].(string)
			stock.Exchange = stockDetail["exchangeName"].(string)
			stock.MarketCap = int64(marketCap["raw"].(float64))
			stock.Price = price["raw"].(float64)
			stock.AverageVolume = int64(volume["raw"].(float64))
			// Company details
			if quoteStore["summaryProfile"] != nil {
				summary := quoteStore["summaryProfile"].(map[string]interface{})
				stock.Sector = summary["sector"].(string)
				stock.Industry = summary["industry"].(string)
				stock.Website = CastToString(summary["website"])
				stock.Country = CastToString(summary["country"])
				stock.Description = CastToString(summary["longBusinessSummary"])
			}

			// News
			streamStore := stores["StreamStore"].(map[string]interface{})
			streams := streamStore["streams"].(map[string]interface{})
			symbolStream := streams[fmt.Sprintf("YFINANCE:%s.mega", symbol)].(map[string]interface{})
			symbolStreamData := symbolStream["data"].(map[string]interface{})
			streamItems := symbolStreamData["stream_items"].([]interface{})

			news := make([]Article, 0, len(streamItems))
			for _, streamItem := range streamItems {
				item := streamItem.(map[string]interface{})
				publishedAt := utils.ConvertUnixMillisToTime(int64(item["pubtime"].(float64)))
				news = append(news, Article{
					URL:         item["url"].(string),
					Title:       item["title"].(string),
					Summary:     item["summary"].(string),
					PublishedAt: publishedAt.In(location),
				})
			}
			stock.News = news

			return stock, err
		}
	}
	return stock, errors.New("could not parse Yahoo Finance response")
}

func CastToString(any interface{}) string {
	if any == nil {
		return ""
	}
	return any.(string)
}

func formatISO(date time.Time) string {
	return date.Format("2006-01-02")
}

func getFloat64(val interface{}) *float64 {
	if val == nil {
		return nil
	}
	value := val.(float64)
	return &value
}

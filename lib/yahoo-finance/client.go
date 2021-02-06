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

const (
	baseURL                              = "https://finance.yahoo.com"
	maxEventsPerPage                     = 100
	TransferAgentSystem EarningsCallTime = "TAS"
	TimeNotSupplied     EarningsCallTime = "TNS"
	BeforeMarketOpen    EarningsCallTime = "BMO"
	AfterMarketClose    EarningsCallTime = "AFC"
)

var parseError = errors.New("failed to parse API response")

type EarningsCallTime string

type EarningsCall struct {
	Ticker             string
	Company            string
	StartTime          EarningsCallTime
	EPSEstimate        *utils.NullFloat64
	EPSActual          *utils.NullFloat64
	EPSSurprisePercent *utils.NullFloat64
	QuoteType          string
	Date               string
}

type IPO struct {
	Ticker    string
	Company   string
	Exchange  string
	PriceFrom *utils.NullFloat64
	PriceTo   *utils.NullFloat64
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
	Sector        *utils.NullString
	Industry      *utils.NullString
	Description   *utils.NullString
	Country       *utils.NullString
	Website       *utils.NullString
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
			if record, ok := row.(map[string]interface{}); ok {
				earningsCall := EarningsCall{
					Ticker:             record["ticker"].(string),
					Company:            record["companyshortname"].(string),
					StartTime:          EarningsCallTime(record["startdatetimetype"].(string)),
					EPSEstimate:        utils.ToNullFloat64(record["epsestimate"]),
					EPSActual:          utils.ToNullFloat64(record["epsactual"]),
					EPSSurprisePercent: utils.ToNullFloat64(record["epssurprisepct"]),
					QuoteType:          record["quoteType"].(string),
					Date:               formattedDate,
				}
				earnings = append(earnings, earningsCall)
			}
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
			if record, ok := row.(map[string]interface{}); ok {
				ipo := IPO{
					Ticker:    record["ticker"].(string),
					Company:   record["companyshortname"].(string),
					Exchange:  record["exchange_short_name"].(string),
					PriceFrom: utils.ToNullFloat64(record["pricefrom"]),
					PriceTo:   utils.ToNullFloat64(record["priceto"]),
					Currency:  record["currencyname"].(string),
					QuoteType: record["quoteType"].(string),
					Date:      formattedDate,
				}
				ipos = append(ipos, ipo)
			}
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
			context, ok := data["context"].(map[string]interface{})
			if !ok {
				return nil, parseError
			}
			dispatcher, ok := context["dispatcher"].(map[string]interface{})
			if !ok {
				return nil, parseError
			}
			stores, ok := dispatcher["stores"].(map[string]interface{})
			if !ok {
				return nil, parseError
			}
			resultsStore, ok := stores["ScreenerResultsStore"].(map[string]interface{})
			if !ok {
				return nil, parseError
			}
			results, ok := resultsStore["results"].(map[string]interface{})
			if !ok {
				return nil, parseError
			}
			rows, ok := results["rows"].([]interface{})
			if !ok {
				return nil, parseError
			}
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
			if marketCap["raw"] != nil {
				stock.MarketCap = int64(marketCap["raw"].(float64))
			}
			stock.Price = price["raw"].(float64)
			stock.AverageVolume = int64(volume["raw"].(float64))
			// Company details
			if quoteStore["summaryProfile"] != nil {
				if summary, ok := quoteStore["summaryProfile"].(map[string]interface{}); ok {
					stock.Sector = utils.ToNullString(summary["sector"])
					stock.Industry = utils.ToNullString(summary["industry"])
					stock.Website = utils.ToNullString(summary["website"])
					stock.Country = utils.ToNullString(summary["country"])
					stock.Description = utils.ToNullString(summary["longBusinessSummary"])
				}
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

func formatISO(date time.Time) string {
	return date.Format("2006-01-02")
}

package yahoo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

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
}

type IPO struct {
	Ticker    string
	Company   string
	Exchange  string
	PriceFrom *float64
	PriceTo   *float64
	Currency  string
	QuoteType string
}

type Client struct {
	client http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: *httpClient,
	}
}

func (c *Client) GetEarningsCall(startDate, endDate string) (earnings []EarningsCall, err error) {
	response, err := c.client.Get("https://finance.yahoo.com/calendar/earnings?from=2020-06-28&to=2020-07-04&day=2020-07-02")
	if err != nil {
		return earnings, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return earnings, err
	}

	bodyString := string(bodyBytes)
	for _, str := range strings.Split(bodyString, "\n") {
		length := len(str)
		if length > 16 && str[:16] == "root.App.main = " {
			var data map[string]interface{}
			if err = json.Unmarshal([]byte(str[16:length-1]), &data); err != nil {
				return earnings, err
			}
			context := data["context"].(map[string]interface{})
			dispatcher := context["dispatcher"].(map[string]interface{})
			stores := dispatcher["stores"].(map[string]interface{})
			resultsStore := stores["ScreenerResultsStore"].(map[string]interface{})
			results := resultsStore["results"].(map[string]interface{})
			rows := results["rows"].([]interface{})
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
				}
				earnings = append(earnings, earningsCall)
			}
			break
		}
	}
	return earnings, nil
}

func (c *Client) GetIPOs(startDate, endDate string) (ipos []IPO, err error) {
	response, err := c.client.Get("https://finance.yahoo.com/calendar/ipo?from=2020-06-28&to=2020-07-04&day=2020-07-02")
	if err != nil {
		return ipos, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ipos, err
	}

	bodyString := string(bodyBytes)
	for _, str := range strings.Split(bodyString, "\n") {
		length := len(str)
		if length > 16 && str[:16] == "root.App.main = " {
			var data map[string]interface{}
			if err = json.Unmarshal([]byte(str[16:length-1]), &data); err != nil {
				return ipos, err
			}
			context := data["context"].(map[string]interface{})
			dispatcher := context["dispatcher"].(map[string]interface{})
			stores := dispatcher["stores"].(map[string]interface{})
			resultsStore := stores["ScreenerResultsStore"].(map[string]interface{})
			results := resultsStore["results"].(map[string]interface{})
			rows := results["rows"].([]interface{})
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
				}
				ipos = append(ipos, ipo)
			}
			break
		}
	}
	return ipos, nil
}

func getFloat64(val interface{}) *float64 {
	if val == nil {
		return nil
	}
	value := val.(float64)
	return &value
}

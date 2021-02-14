package alpaca

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const liveAPIPath = "https://api.alpaca.markets/v2"
const paperAPIPath = "https://paper-api.alpaca.markets/v2"

// Data API endpoints still use v1
// https://alpaca.markets/docs/api-documentation/api-v2/market-data/
const marketDataApiPath = "https://data.alpaca.markets/v1"

type Client struct {
	client       http.Client
	apiKey       string
	apiSecretKey string
	basePath     string
}

// Alpaca data is currently limited to 5 US exchanges compared to Polygon which
// consolidates from all exchanges in the US
// https://alpaca.markets/docs/api-documentation/api-v2/market-data/#which-api-should-i-use
func NewClient(httpClient *http.Client, apiKey string, apiSecretKey string, isLive bool) *Client {
	basePath := paperAPIPath
	if isLive {
		basePath = liveAPIPath
	}
	return &Client{
		client:       *httpClient,
		apiKey:       apiKey,
		apiSecretKey: apiSecretKey,
		basePath:     basePath,
	}
}

func (c *Client) GetAssets(status, assetClass string) (assets []Asset, err error) {
	req, err := http.NewRequest("GET", c.basePath+"/assets", nil)
	if err != nil {
		return assets, err
	}
	c.setHeaders(req)
	queryParams := req.URL.Query()
	if status != "" {
		queryParams.Add("status", status)
	}
	if assetClass != "" {
		queryParams.Add("asset_class", assetClass)
	}
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return assets, err
	}
	if resp.StatusCode != http.StatusOK {
		return assets, errors.New("Response failed with status code: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&assets); err != nil {
		return assets, err
	}
	return assets, nil
}

type CandleQueryParams struct {
	Symbols    []string
	Limit      int
	CandleSize CandleSize
	StartTime  time.Time
	EndTime    time.Time
}

// NOTE The number of candles returned is controlled by params.Limit
func (c *Client) GetCandlesBySymbol(params CandleQueryParams) (candles map[string][]Candle, err error) {
	if len(params.Symbols) == 0 || len(params.Symbols) > 200 {
		return candles, errors.New("symbols must be between 1 to 200")
	}
	if params.Limit > 1000 {
		return candles, errors.New("limit must be between 1 to 1000")
	}
	switch params.CandleSize {
	case OneMin, FiveMin, FifteenMin, OneDay:
		break
	default:
		return candles, errors.New("candleSize must be supported CandleSize")
	}

	url := marketDataApiPath + "/bars/" + string(params.CandleSize)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return candles, err
	}

	c.setHeaders(req)
	queryParams := req.URL.Query()
	queryParams.Add("symbols", strings.Join(params.Symbols, ","))
	if params.Limit > 0 {
		queryParams.Add("limit", strconv.Itoa(params.Limit))
	}
	if !params.StartTime.IsZero() {
		queryParams.Add("start", formatTime(params.StartTime))
	}
	if !params.EndTime.IsZero() {
		queryParams.Add("end", formatTime(params.EndTime))
	}
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return candles, err
	}
	if resp.StatusCode != http.StatusOK {
		return candles, errors.New("Response failed with status code: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&candles); err != nil {
		return candles, err
	}
	return candles, nil
}

type TradeResponse struct {
	Status string    `json:"status"`
	Symbol string    `json:"symbol"`
	Last   LastTrade `json:"last"`
}

func (c *Client) GetLastTrade(symbol string) (trade LastTrade, err error) {
	url := marketDataApiPath + "/last/stocks/" + symbol
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return trade, err
	}

	c.setHeaders(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return trade, err
	}
	if resp.StatusCode != http.StatusOK {
		return trade, errors.New("Response failed with status code: " + resp.Status)
	}
	var tradeResp TradeResponse
	if err := json.NewDecoder(resp.Body).Decode(&tradeResp); err != nil {
		return trade, err
	}
	return tradeResp.Last, err
}

type QuoteResponse struct {
	Status string    `json:"status"`
	Symbol string    `json:"symbol"`
	Last   LastQuote `json:"last"`
}

func (c *Client) GetLastQuote(symbol string) (quote LastQuote, err error) {
	url := marketDataApiPath + "/last_quote/stocks/" + symbol
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return quote, err
	}

	c.setHeaders(req)
	resp, err := c.client.Do(req)
	if err != nil {
		return quote, err
	}
	if resp.StatusCode != http.StatusOK {
		return quote, errors.New("Response failed with status code: " + resp.Status)
	}
	var quoteResp QuoteResponse
	if err := json.NewDecoder(resp.Body).Decode(&quoteResp); err != nil {
		return quote, err
	}
	return quoteResp.Last, err
}

func (c *Client) setHeaders(request *http.Request) {
	request.Header.Set("APCA-API-KEY-ID", c.apiKey)
	request.Header.Set("APCA-API-SECRET-KEY", c.apiSecretKey)
}

func formatTime(date time.Time) string {
	return date.Format("2006-01-02T15:04:05-07:00")
}

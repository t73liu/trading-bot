package alpaca

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

const liveAPIPath = "https://api.alpaca.markets/v2"
const paperAPIPath = "https://paper-api.alpaca.markets/v2"

// https://alpaca.markets/docs/api-documentation/api-v2/market-data/
const marketDataApiPath = "https://data.alpaca.markets/v2"

type Client struct {
	config   ClientConfig
	basePath string
	wsConn   *websocket.Conn
}

type ClientConfig struct {
	HttpClient    *http.Client
	ApiKey        string
	ApiSecret     string
	IsLiveTrading bool
	IsPaidData    bool
}

// Alpaca's free plan only provides data from IEX whereas the paid plan provides
// data from all US exchanges
// https://alpaca.markets/docs/api-documentation/api-v2/market-data/alpaca-data-api-v2/#subscription-plans
func NewClient(config ClientConfig) *Client {
	basePath := paperAPIPath
	if config.IsLiveTrading {
		basePath = liveAPIPath
	}
	client := &Client{
		config:   config,
		basePath: basePath,
	}
	return client
}

func (c *Client) Close() error {
	if c.wsConn != nil {
		return c.wsConn.Close()
	}
	return nil
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

	resp, err := c.config.HttpClient.Do(req)
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
	Limit      int
	CandleSize CandleSize
	StartTime  time.Time
	EndTime    time.Time
	PageToken  string
}

type CandlesResponse struct {
	Candles       []Candle `json:"bars"`
	Symbol        string   `json:"symbol"`
	NextPageToken string   `json:"next_page_token"`
}

func (c *Client) GetSymbolCandles(symbol string, params CandleQueryParams) (candlesResponse CandlesResponse, err error) {
	if params.Limit > 10000 {
		return candlesResponse, errors.New("limit must be between 1 to 10000")
	}
	switch params.CandleSize {
	case OneMin, OneHour, OneDay:
		break
	default:
		return candlesResponse, errors.New("candleSize must be supported CandleSize")
	}

	url := fmt.Sprintf("%s/stocks/%s/bars", marketDataApiPath, symbol)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return candlesResponse, err
	}

	c.setHeaders(req)
	queryParams := req.URL.Query()
	queryParams.Add("timeframe", string(params.CandleSize))
	if params.Limit > 0 {
		queryParams.Add("limit", strconv.Itoa(params.Limit))
	}
	if !params.StartTime.IsZero() {
		queryParams.Add("start", formatTime(params.StartTime))
	}
	if !params.EndTime.IsZero() {
		queryParams.Add("end", formatTime(params.EndTime))
	}
	if params.PageToken != "" {
		queryParams.Add("page_token", params.PageToken)
	}
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.config.HttpClient.Do(req)
	if err != nil {
		return candlesResponse, err
	}
	if resp.StatusCode != http.StatusOK {
		return candlesResponse, errors.New("Response failed with status code: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&candlesResponse); err != nil {
		return candlesResponse, err
	}
	return candlesResponse, nil
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
	resp, err := c.config.HttpClient.Do(req)
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
	resp, err := c.config.HttpClient.Do(req)
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
	request.Header.Set("APCA-API-KEY-ID", c.config.ApiKey)
	request.Header.Set("APCA-API-SECRET-KEY", c.config.ApiSecret)
}

func formatTime(date time.Time) string {
	return date.UTC().Format("2006-01-02T15:04:05Z")
}

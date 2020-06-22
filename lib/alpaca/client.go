package alpaca

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
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

func (c *Client) GetCandles(candleSize CandleSize, symbols []string, limit int) (candles map[string][]Candle, err error) {
	if len(symbols) == 0 || len(symbols) > 200 {
		return candles, errors.New("symbols must be between 1 to 200")
	}
	if limit > 1000 {
		return candles, errors.New("limit must be between 1 to 1000")
	}
	switch candleSize {
	case OneMin, FiveMin, FifteenMin, OneDay:
		break
	default:
		return candles, errors.New("candleSize must be supported CandleSize")
	}

	req, err := http.NewRequest("GET", marketDataApiPath+"/bars"+string(candleSize), nil)
	if err != nil {
		return candles, err
	}
	c.setHeaders(req)
	queryParams := req.URL.Query()
	queryParams.Add("symbols", strings.Join(symbols, ","))
	if limit != 0 {
		queryParams.Add("limit", strconv.Itoa(limit))
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

func (c *Client) setHeaders(request *http.Request) {
	request.Header.Set("APCA-API-KEY-ID", c.apiKey)
	request.Header.Set("APCA-API-SECRET-KEY", c.apiSecretKey)
}

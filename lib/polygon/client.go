package polygon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const polygonHost = "https://api.polygon.io"

type Client struct {
	client http.Client
	apiKey string
}

func NewClient(httpClient *http.Client, apiKey string) *Client {
	return &Client{
		client: *httpClient,
		apiKey: apiKey,
	}
}

type TickersQueryParams struct {
	PerPage int
	Page    int
	Active  bool
	Type    string
	Market  string
	Locale  string
	Search  string
	Sort    string
}

func DefaultTickersQueryParams() TickersQueryParams {
	return TickersQueryParams{
		PerPage: 50,
		Page:    1,
		Active:  true,
		Type:    "",
		Market:  "stocks",
		Locale:  "us",
		Search:  "",
		Sort:    "ticker",
	}
}

func (c *Client) GetTickers(params TickersQueryParams) (result TickersResponse, err error) {
	req, err := http.NewRequest("GET", polygonHost+"/v2/reference/tickers", nil)
	if err != nil {
		return result, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("apiKey", c.apiKey)
	if params.Sort != "" {
		queryParams.Add("sort", params.Sort)
	}
	if params.Type != "" {
		queryParams.Add("type", params.Type)
	}
	if params.Market != "" {
		queryParams.Add("market", params.Market)
	}
	if params.Locale != "" {
		queryParams.Add("locale", params.Locale)
	}
	if params.Search != "" {
		queryParams.Add("search", params.Search)
	}
	if params.PerPage > 0 {
		queryParams.Add("perpage", strconv.Itoa(params.PerPage))
	}
	if params.Page > 0 {
		queryParams.Add("page", strconv.Itoa(params.Page))
	}
	queryParams.Add("active", strconv.FormatBool(params.Active))
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return result, err
	}
	if resp.StatusCode != http.StatusOK {
		return result, errors.New("Response failed with status code: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, err
	}
	return result, nil
}

// News articles have stopped updating since 2020-03-26
// https://github.com/polygon-io/issues/issues/25
func (c *Client) GetTickerNews(ticker string, perPage int, page int) (articles []Article, err error) {
	url := polygonHost + "/v1/meta/symbols/" + ticker + "/news"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return articles, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("apiKey", c.apiKey)
	queryParams.Add("perpage", strconv.Itoa(perPage))
	queryParams.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return articles, err
	}
	if resp.StatusCode != http.StatusOK {
		return articles, errors.New("Response failed with status code: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&articles); err != nil {
		return articles, err
	}
	return articles, nil
}

// Some tickers return 404 despite being listed as active with bars
// https://github.com/polygon-io/issues/issues/40
func (c *Client) GetTickerDetails(ticker string) (detail TickerDetails, err error) {
	url := polygonHost + "/v1/meta/symbols/" + ticker + "/company"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return detail, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("apiKey", c.apiKey)
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return detail, err
	}
	if resp.StatusCode != http.StatusOK {
		return detail, errors.New("Response failed with status code: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&detail); err != nil {
		return detail, err
	}
	return detail, nil
}

type TickerBarsQueryParams struct {
	Ticker       string
	TimeInterval int
	TimeUnit     string
	StartDate    time.Time
	EndDate      time.Time
	Unadjusted   bool
	Sort         string
}

// Maxed number of bars returned is 5000
// https://github.com/polygon-io/issues/issues/45
func (c *Client) GetTickerBars(params TickerBarsQueryParams) (bars []TickerBar, err error) {
	if params.TimeInterval < 1 {
		return bars, errors.New("TimeInterval must be at least 1")
	}
	if !params.EndDate.After(params.StartDate) {
		return bars, errors.New("EndDate must be after StartDate")
	}
	url := fmt.Sprintf(
		"%s/v2/aggs/ticker/%s/range/%d/%s/%s/%s",
		polygonHost,
		params.Ticker,
		params.TimeInterval,
		params.TimeUnit,
		formatDate(params.StartDate),
		formatDate(params.EndDate),
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return bars, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("apiKey", c.apiKey)
	if params.Unadjusted {
		queryParams.Add("unadjusted", "true")
	}
	if params.Sort != "" {
		queryParams.Add("sort", params.Sort)
	}
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return bars, err
	}
	if resp.StatusCode != http.StatusOK {
		return bars, errors.New("Response failed with status code: " + resp.Status)
	}
	var barsResponse TickerBarsResponse
	if err := json.NewDecoder(resp.Body).Decode(&barsResponse); err != nil {
		return bars, err
	}
	return barsResponse.Results, nil
}

func formatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

type TickerSnapshotResponse struct {
	Status  string           `json:"status"`
	Tickers []TickerSnapshot `json:"tickers"`
}

func (c *Client) GetMovers(isIncreasing bool) (snapshots []TickerSnapshot, err error) {
	path := "gainers"
	if !isIncreasing {
		path = "losers"
	}
	url := fmt.Sprintf(
		"%s/v2/snapshot/locale/us/markets/stocks/%s",
		polygonHost,
		path,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return snapshots, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("apiKey", c.apiKey)
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return snapshots, err
	}
	if resp.StatusCode != http.StatusOK {
		return snapshots, errors.New("Response failed with status code: " + resp.Status)
	}
	var snapshotsResponse TickerSnapshotResponse
	if err := json.NewDecoder(resp.Body).Decode(&snapshotsResponse); err != nil {
		return snapshots, err
	}
	return snapshotsResponse.Tickers, nil
}

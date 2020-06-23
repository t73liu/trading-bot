package polygon

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

type TickersResponse struct {
	Count   int      `json:"count"`
	Tickers []Ticker `json:"tickers"`
}

type Ticker struct {
	Ticker          string `json:"ticker"`
	Name            string `json:"name"`
	Market          string `json:"market"`
	Locale          string `json:"locale"`
	Currency        string `json:"currency"`
	Active          bool   `json:"active"`
	PrimaryExchange string `json:"primaryExch"`
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

func (c *Client) GetTickers(params TickersQueryParams) (result TickersResponse, err error) {
	req, err := http.NewRequest("GET", polygonHost+"/v2/reference/tickers", nil)
	if err != nil {
		return result, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("apiKey", c.apiKey)
	queryParams.Add("sort", params.Sort)
	queryParams.Add("type", params.Type)
	queryParams.Add("market", params.Market)
	queryParams.Add("locale", params.Locale)
	queryParams.Add("search", params.Search)
	queryParams.Add("perpage", strconv.Itoa(params.PerPage))
	queryParams.Add("page", strconv.Itoa(params.Page))
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

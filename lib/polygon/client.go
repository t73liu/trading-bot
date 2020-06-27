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

// TODO Implement GetTickerBars
func (c *Client) GetTickerBars(ticker string) (bars interface{}, err error) {
	url := polygonHost + "/v2/aggs/ticker/" + ticker + "/range/1/hour/2020-06-25/2020-06-26"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return bars, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("apiKey", c.apiKey)
	queryParams.Add("unadjusted", "true")
	queryParams.Add("sort", "asc")
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return bars, err
	}
	if resp.StatusCode != http.StatusOK {
		return bars, errors.New("Response failed with status code: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&bars); err != nil {
		return bars, err
	}
	return bars, nil
}

package iex

import (
	"encoding/json"
	"errors"
	"net/http"
)

const baseURL = "https://cloud.iexapis.com/stable"

type Client struct {
	client http.Client
	token  string
}

func NewClient(httpClient *http.Client, token string) *Client {
	return &Client{
		client: *httpClient,
		token:  token,
	}
}

type Stock struct {
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

func (c *Client) GetReferenceSymbols() (stocks []Stock, err error) {
	url := baseURL + "/ref-data/symbols"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	queryParams := req.URL.Query()
	queryParams.Add("token", c.token)
	req.URL.RawQuery = queryParams.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Response failed with status code: " + resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&stocks); err != nil {
		return nil, err
	}
	return stocks, nil
}

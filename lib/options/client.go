package options

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const optionsURL = "https://www.optionsprofitcalculator.com/ajax/getOptions?stock=%s&reqId=1"

type Client struct {
	client http.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: *httpClient,
	}
}

type Response struct {
	Options map[string]interface{} `json:"options"`
}

type Option struct {
	Type           string  `json:"type"`
	ExpirationDate string  `json:"expirationDate"`
	StrikePrice    float64 `json:"strikePrice"`
	Last           float64 `json:"lastPrice"`
	Bid            float64 `json:"bid"`
	Ask            float64 `json:"ask"`
	Volume         int64   `json:"volume"`
}

// Options are 15 minute delayed
func (c *Client) GetOptions(symbol string) (options []Option, err error) {
	resp, err := http.Get(fmt.Sprintf(optionsURL, symbol))
	if err != nil {
		return options, err
	}
	var result Response
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return options, err
	}
	if len(result.Options) == 0 {
		return options, nil
	}
	for key, val := range result.Options {
		if key == "_data_source" {
			continue
		}
		optionsByTypeAndStrike := val.(map[string]interface{})
		calls := optionsByTypeAndStrike["c"].(map[string]interface{})
		for strike, optionDetail := range calls {
			option, err := parseOptions("call", key, strike, optionDetail.(map[string]interface{}))
			if err != nil {
				return options, err
			}
			options = append(options, option)
		}
		puts := optionsByTypeAndStrike["p"].(map[string]interface{})
		for strike, optionDetail := range puts {
			option, err := parseOptions("put", key, strike, optionDetail.(map[string]interface{}))
			if err != nil {
				return options, err
			}
			options = append(options, option)
		}
	}
	return options, nil
}

func parseOptions(optionType string, expiry string, strike string, optionDetail map[string]interface{}) (option Option, err error) {
	strikePrice, err := strconv.ParseFloat(strike, 64)
	if err != nil {
		return option, err
	}
	last, err := strconv.ParseFloat(optionDetail["l"].(string), 64)
	if err != nil {
		return option, err
	}
	bid, err := strconv.ParseFloat(optionDetail["b"].(string), 64)
	if err != nil {
		return option, err
	}
	ask, err := strconv.ParseFloat(optionDetail["a"].(string), 64)
	if err != nil {
		return option, err
	}
	volume, err := strconv.Atoi(optionDetail["v"].(string))
	if err != nil {
		return option, err
	}
	option = Option{
		Type:           optionType,
		ExpirationDate: expiry,
		StrikePrice:    strikePrice,
		Volume:         int64(volume),
		Bid:            bid,
		Ask:            ask,
		Last:           last,
	}
	return option, nil
}

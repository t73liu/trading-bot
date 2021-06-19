## Alpaca API

Go client for [Alpaca v2](https://alpaca.markets/).

## Usage

```golang
package main

import (
	"github.com/t73liu/trading-bot/lib/alpaca"
)

func main() {
	httpClient := &http.Client{Timeout: 15 * time.Second}
	alpacaClient := alpaca.NewClient(alpaca.ClientConfig{
		HttpClient:    &httpClient,
		ApiKey:        "API_KEY",
		ApiSecret:     "API_SECRET",
		IsLiveTrading: false,
		IsPaidData:    false,
	})
	assets, err := alpacaClient.GetAssets("active", "us_equities")
}
```

## Alpaca API

Go client for [Alpaca](https://alpaca.markets/).

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/alpaca"

func main() {
  httpClient := &http.Client{Timeout: 15 * time.Second}
  alpacaClient := alpaca.Client(httpClient, "API_KEY")
  alpacaClient.GetAssets("active", "us_equities")
}
```

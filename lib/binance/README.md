# Binance API

Go client for [Binance](https://www.binance.com/en).

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/binance"

func main() {
  httpClient := &http.Client{Timeout: 15 * time.Second}
  binanceClient := binance.Client(httpClient, "API_KEY")
  binanceClient.GetExchangeInfo()
}
```

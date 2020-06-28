## Polygon API

Go client for [Polygon](https://polygon.io/).

Refer to the following repo to track ongoing issues: https://github.com/polygon-io/issues/issues

**Note:** API key is provided by Alpaca.

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/polygon"

func main() {
  httpClient := &http.Client{Timeout: 15 * time.Second}
  polygonClient := polygon.Client(httpClient, "API_KEY")
  polygonClient.GetTickers(polygon.TickerQueryParams{})
}
```

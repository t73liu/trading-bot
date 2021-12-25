# FinViz API

Scrape FinViz overview table by providing query string.

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/finviz"

func main() {
  httpClient := &http.Client{Timeout: 15 * time.Second}
  finvizClient := finviz.Client(httpClient)
  // Ordered by descending gap with lower limit of +5%
  finvizClient.ScreenStocks("v=111&f=ta_gap_u5&ft=4&o=-gap")
}
```

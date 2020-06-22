## News API

Go client for [News API](https://newsapi.org/).

**NOTE:** "Developer" plan has one hour delay.

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/newsapi"

func main() {
  httpClient := &http.Client{Timeout: 15 * time.Second}
  newsClient := newsapi.Client(httpClient, "API_KEY")
  newsClient.GetTopHeadlinesBySources("bitcoin", "the-wall-street-journal")
}
```

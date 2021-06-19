## Trader DB

Go module containing common `traderdb` queries and models.

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/traderdb"

func main() {
	dbPool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		logger.Fatalln("Unable to connect to database:", err)
	}
	defer dbPool.Close()
	watchlists, err := traderdb.GetWatchlistsWithUserID(dbPool, 1)
	// ...
}
```

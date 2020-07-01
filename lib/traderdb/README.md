## Trader DB

Go module containing common `traderdb` queries and models.

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/traderdb"

func main() {
	dbPool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Fatalln("Unable to connect to database:", err)
	}
	defer dbPool.Close()
	watchlists, err := traderdb.GetWatchlistsByUserId(dbPool, 1)
	// ...
}
```

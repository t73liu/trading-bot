# Trader

`trader` is a Go server responsible for the following tasks:

- Serving UI static assets in production
- Handling UI API calls
- Updating DB with the latest stock/crypto trades

## Prerequisites

The DB needs to be running before starting `trader`. Instructions for setting up
the DB can be found in the [README.md](../traderdb/README.md).

`trader` can be run without the UI (Dash) by building Dash and serving its static
assets. This only needs to be done once.
 
```bash
cd ${TRADING_BOT_REPO}/dash
yarn build
mv ${TRADING_BOT_REPO}/dash/build ${TRADING_BOT_REPO}/trader/assets
```

## Development

`trader` can be configured with the following flags:

- `-db.url` (e.g. `postgres://postgres:test@localhost:5432/traderdb?sslmode=disable`)
- `-alpaca.key`: Use the paper trading API key during development
- `-alpaca.secret`: Use the paper trading API secret during development
- `-news.key`

```bash
cd ${TRADING_BOT_REPO}/trader
go run . -db.url DB_URL -alpaca.key API_KEY -alpaca.secret API_SECRET
```

## Rate Limits

- Alpaca: 200 per minute for each api key (live/paper)
- Binance: 1200 per minute and 100 orders per 10 seconds
- NewsAPI: 500 per day

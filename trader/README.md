## Trader

Trader is a Go server responsible for the following tasks:

- Serving UI static assets in production
- Handling UI API calls
- Updating DB with latest stock/crypto info

### Prerequisites

The DB needs to be running before starting Trader. Instructions for setting up
the DB can be found in the [README.md](/db/README.md).

Trader can be ran without the UI (Dash) by building Dash and serving its static
assets. This only needs to be done once.
 
```sh
cd ${TRADING_BOT_REPO}/dash
yarn build
mv ${TRADING_BOT_REPO}/dash/build ${TRADING_BOT_REPO}/trader/assets
```

### Development

Trader can be configured with the following environment variables:

- `DB_URL`
- `DB_USER`
- `ALPACA_API_KEY` (optional): Use the paper trading API key during development
- `ALPACA_API_SECRET` (optional): Use the paper trading API secret during development
- `BINANCE_API_KEY` (optional)
- `BINANCE_API_SECRET` (optional)
- `NEWS_API_KEY` (optional)

```sh
cd ${TRADING_BOT_REPO}/trader
go run main.go
```

### Rate Limits

- Alpaca: 200 per minute per api key (live/paper)
- Binance: 1200 per minute and 100 orders per 10 seconds
- NewsAPI: 500 per day

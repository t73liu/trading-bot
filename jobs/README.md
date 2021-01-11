## Jobs

This directory contains periodically scheduled jobs.

- `earnings-ipo-news`: Send weekly email for Earnings/IPOs in the upcoming 2 weeks
- `populate-stocks`: Check Alpaca daily for new stocks and updated stock info
    - `DATABASE_URL` (e.g. `postgres://postgres:test@localhost:5432/traderdb?sslmode=disable`)
    - `ALPACA_API_KEY`: Use the paper trading API key during development
    - `ALPACA_API_SECRET`: Use the paper trading API secret during development
    - `IEX_API_TOKEN`: Used to fetch proper company names (costs 100 message per call)
- `watchlist-news`: Send daily email for watchlist stocks since last trading day

## Trading Bot

### Features (WIP)

- [x] Fetch news from [News API](https://newsapi.org/)
- [ ] Fetch stock data from [Polygon](https://polygon.io/)
- [ ] Live and paper stock trading with [Alpaca](https://alpaca.markets/)
- [ ] Live crypto trading with [Binance](https://www.binance.com/en)

**Note:** Functionality requires setting up API keys with corresponding services.
Polygon API key is provided by Alpaca.

### Getting Started

This repository is divided into the following sections:

- `adhoc`: Useful ad hoc scripts and experiments
- `dash`: UI for visualizing trades and account balances
- `db`: Database containing financial, ML and user info
- `jobs`: Periodically scheduled jobs
- `lib`: Shared Go libraries
- `quant`: Machine learning models and trading strategies
- `trader`: Backend server for UI and makes periodic trades

This directory will be referred to by the `TRADING_BOT_REPO` environment variable.

### Technologies

- [Go 1.14](https://golang.org/)
- [Node 12](https://nodejs.org/en/)
- [PostgreSQL 12](https://www.postgresql.org/)
- [Python 3.7](https://www.python.org/)
- [React 16.13](https://reactjs.org/)
- [TensorFlow 2](https://www.tensorflow.org/)
- [Terraform 0.12](https://www.terraform.io/)
- [Yarn 1.21](https://classic.yarnpkg.com/lang/en/)

## Trading Bot

### Features (WIP)

- [x] Fetch news from News API
- [ ] Fetch stock data from Polygon
- [ ] Live and paper stock trading with Alpaca
- [ ] Live crypto trading with Binance

Note: Functionality requires setting up API keys with corresponding services.
Polygon API key is provided by Alpaca.

### Getting Started

This repository is divided into the following sections:

- `adhoc`: Useful ad hoc scripts and experiments
- `dash`: UI for visualizing trades and account balances
- `db`: Database containing financial, ML and user info
- `quant`: Machine learning models and trading strategies
- `trader`: Backend server for UI and makes periodic trades

This directory will be referred to by the `TRADING_BOT_REPO` environment variable.

### Technologies

- Go 1.14
- Node 12
- Postgres 12
- Python 3.7
- React 16.13
- TensorFlow 2
- Terraform 0.12
- Yarn 1.21

### Future Ideas

- Send Email/SMS alerts via Twilio

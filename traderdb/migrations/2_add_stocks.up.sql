CREATE TABLE IF NOT EXISTS stocks (
  id SERIAL PRIMARY KEY,
  symbol TEXT UNIQUE NOT NULL,
  company TEXT NOT NULL,
  exchange TEXT NOT NULL,
  tradable BOOLEAN NOT NULL DEFAULT TRUE,
  marginable BOOLEAN NOT NULL DEFAULT FALSE,
  shortable BOOLEAN NOT NULL DEFAULT FALSE,
  price_micros BIGINT,
  market_cap BIGINT,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS stock_candles (
  stock_id INTEGER REFERENCES stocks(id) NOT NULL,
  opened_at TIMESTAMPTZ NOT NULL,
  open_micros BIGINT NOT NULL,
  high_micros BIGINT NOT NULL,
  low_micros BIGINT NOT NULL,
  close_micros BIGINT NOT NULL,
  volume INTEGER NOT NULL,

  PRIMARY KEY (stock_id, opened_at)
);

CREATE TABLE IF NOT EXISTS stock_positions (
  id SERIAL PRIMARY KEY,
  stock_id INTEGER REFERENCES stocks(id) NOT NULL,
  user_id INTEGER REFERENCES users(id) NOT NULL,
  entry_price_micros BIGINT NOT NULL,
  number_of_shares INTEGER NOT NULL,
  purchased_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

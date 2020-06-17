CREATE TABLE IF NOT EXISTS watchlists (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) NOT NULL,
  name text NOT NULL
);

CREATE TABLE IF NOT EXISTS watchlist_stocks (
  watchlist_id INTEGER REFERENCES watchlists(id) NOT NULL,
  stock_id INTEGER REFERENCES stocks(id) NOT NULL,

  PRIMARY KEY (watchlist_id, stock_id)
);

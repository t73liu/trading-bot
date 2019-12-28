CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    email text UNIQUE NOT NULL,
    password text NOT NULL
);

CREATE TABLE IF NOT EXISTS stocks (
    id serial PRIMARY KEY,
    symbol text UNIQUE NOT NULL,
    company text NOT NULL,
    price NUMERIC (9,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS stock_candles (
    id serial PRIMARY KEY,
    stock_id INTEGER REFERENCES stocks(id),
    open NUMERIC (9,2) NOT NULL,
    high NUMERIC (9,2) NOT NULL,
    low NUMERIC (9,2) NOT NULL,
    close NUMERIC (9,2) NOT NULL,
    volume INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS cryptos (
    id serial PRIMARY KEY,
    symbol text UNIQUE NOT NULL,
    name text NOT NULL,
    price NUMERIC (13,6) NOT NULL
);

CREATE TABLE IF NOT EXISTS crypto_candles (
    id serial PRIMARY KEY,
    crypto_id INTEGER REFERENCES cryptos(id),
    open NUMERIC (13,6) NOT NULL,
    high NUMERIC (13,6) NOT NULL,
    low NUMERIC (13,6) NOT NULL,
    close NUMERIC (13,6) NOT NULL,
    volume INTEGER NOT NULL
);

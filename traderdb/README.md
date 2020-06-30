## Trader DB

This directory contains migrations for a PostgreSQL database called `traderdb`.
This database contains financial information (e.g. OHLC, quarterly reports) and
user information (e.g. watchlist, positions).

### Prerequisites

Migrations are done with the [migrate](https://github.com/golang-migrate/migrate)
CLI. Usage instructions can be found [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

```sh
# Example migrating local DB with user = "postgres" and password = "test"
cd ${TRADING_BOT_REPO}
migrate -path traderdb/migrations \
 -database "postgres://postgres:test@localhost:5432/traderdb?sslmode=disable" \
 up 1

# Fix failed migration manually and force correct migration version number (e.g. 4)
migrate -path traderdb/migrations \
 -database "postgres://postgres:test@localhost:5432/traderdb?sslmode=disable" \
 force 4
```

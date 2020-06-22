## DB

This directory contains migrations for a PostgreSQL database. This database ("trader")
contains financial information (e.g. OHLC, quarterly reports) and user information
(e.g. watchlist, positions).

### Prerequisites

Migrations are done with the [migrate](https://github.com/golang-migrate/migrate)
CLI. Usage instructions can be found [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

```sh
# Example migrating local DB with user = "postgres" and password = "test"
cd ${TRADING_BOT_REPO}
migrate -path db/migrations \
 -database "postgres://postgres:test@localhost:5432/trader?sslmode=disable" \
 up 1

# Fix failed migration manually and force correct migration version number (e.g. 4)
migrate -path db/migrations \
 -database "postgres://postgres:test@localhost:5432/trader?sslmode=disable" \
 force 4
```

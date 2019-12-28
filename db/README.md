## DB

This directory contains migrations for a PostgreSQL database.
This database contains financial information such as OHLC and
quarterly reports.

### Prerequisites

Migrations are done with the [migrate](https://github.com/golang-migrate/migrate)
CLI. Installation instructions can be found [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate). 

```sh
migrate -database DB_URL -path ${TRADING_BOT_REPO}/db/migrations up
```

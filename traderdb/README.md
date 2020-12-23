## Trader DB

This directory contains migrations for a PostgreSQL database called `traderdb`.
This database contains financial information (e.g. OHLC, quarterly reports) and
user information (e.g. watchlist, positions).

### Development

Download [Docker](https://www.docker.com/) and run the following commands in order
to bring up a PostgreSQL database.

```shell
# Pull the PostgreSQL docker image
docker pull postgres:12-alpine

# Run the PostgreSQL docker
docker run --detach \
 --name traderdb \
 --env POSTGRES_PASSWORD=test \
 --volume $HOME/pgdata/:/var/lib/postgresql/data \
 --publish 5432:5432 \
 postgres:12-alpine
 
# Create PostgreSQL DB
docker exec -it traderdb createdb --username=postgres traderdb

# Access PSQL CLI
docker exec -it traderdb psql --username=postgres traderdb
```

### Migrations

Migrations are done with the [migrate](https://github.com/golang-migrate/migrate)
CLI. Usage instructions can be found [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate).

```shell
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

module trading-bot/adhoc/populate-db

go 1.14

replace github.com/t73liu/trading-bot/lib/alpaca => ../../lib/alpaca

replace github.com/t73liu/trading-bot/lib/newsapi => ../../lib/newsapi

replace github.com/t73liu/trading-bot/lib/polygon => ../../lib/polygon

require (
	github.com/jackc/pgx/v4 v4.7.0
	github.com/t73liu/trading-bot/lib/alpaca v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/newsapi v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/polygon v0.0.0-00010101000000-000000000000
)

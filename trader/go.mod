module github.com/t73liu/trading-bot/trader

go 1.14

replace github.com/t73liu/trading-bot/lib/newsapi => ../lib/newsapi

replace github.com/t73liu/trading-bot/lib/traderdb => ../lib/traderdb

require (
	github.com/caddyserver/certmagic v0.11.2
	github.com/jackc/pgx/v4 v4.7.1
	github.com/julienschmidt/httprouter v1.3.0
	github.com/t73liu/trading-bot/lib/newsapi v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/traderdb v0.0.0-00010101000000-000000000000
)

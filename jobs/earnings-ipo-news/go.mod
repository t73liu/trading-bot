module github.com/t73liu/trading-bot/jobs/earnings-ipo-news

go 1.14

replace github.com/t73liu/trading-bot/lib/traderdb => ../../lib/traderdb

replace github.com/t73liu/trading-bot/lib/yahoo-finance => ./../../lib/yahoo-finance

require (
	github.com/jackc/pgx/v4 v4.7.1
	github.com/t73liu/trading-bot/lib/traderdb v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/yahoo-finance v0.0.0-00010101000000-000000000000
)

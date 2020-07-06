module trading-bot/adhoc/test-strategy

go 1.14

replace github.com/t73liu/trading-bot/lib/technical-analysis => ../../lib/technical-analysis

replace github.com/t73liu/trading-bot/lib/traderdb => ../../lib/traderdb

require (
	github.com/jackc/pgx/v4 v4.7.1
	github.com/t73liu/trading-bot/lib/technical-analysis v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/traderdb v0.0.0-00010101000000-000000000000
)

module trading-bot/adhoc/gap

go 1.14

replace github.com/t73liu/trading-bot/lib/finviz => ../../lib/finviz

replace github.com/t73liu/trading-bot/lib/newsapi => ../../lib/newsapi

replace github.com/t73liu/trading-bot/lib/polygon => ../../lib/polygon

replace github.com/t73liu/trading-bot/lib/traderdb => ../../lib/traderdb

replace github.com/t73liu/trading-bot/lib/utils => ../../lib/utils

require (
	github.com/jackc/pgx/v4 v4.7.1
	github.com/t73liu/trading-bot/lib/finviz v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/newsapi v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/polygon v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/traderdb v0.0.0-00010101000000-000000000000
	github.com/t73liu/trading-bot/lib/utils v0.0.0-00010101000000-000000000000
)

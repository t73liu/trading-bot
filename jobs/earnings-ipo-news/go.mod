module github.com/t73liu/trading-bot/jobs/earnings-ipo-news

go 1.14

replace github.com/t73liu/trading-bot/lib/yahoo-stock-calendar => ../../lib/yahoo-stock-calendar

require (
	github.com/t73liu/trading-bot/lib/yahoo-stock-calendar v0.0.0-00010101000000-000000000000
)

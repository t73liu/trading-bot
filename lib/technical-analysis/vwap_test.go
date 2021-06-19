package analyze

import (
	"testing"

	"github.com/t73liu/tradingbot/lib/candle"
	"github.com/t73liu/tradingbot/lib/utils"
)

func TestVWAP(t *testing.T) {
	t.Run(
		"VWAP empty candles returns nil",
		testVWAPFunc([]candle.Candle{}, nil),
	)
	t.Run(
		"VWAP first value is average of high, low, close",
		testVWAPFunc(
			[]candle.Candle{
				{
					Volume:      89329,
					HighMicros:  utils.DollarsToMicros(127.36),
					LowMicros:   utils.DollarsToMicros(126.99),
					CloseMicros: utils.DollarsToMicros(127.28),
				},
			},
			// (127.36 + 126.99 + 127.28) / 3 = 11
			[]utils.MicroDollar{utils.NewMicroDollar(utils.DollarsToMicros(127.21))},
		),
	)
	t.Run(
		"VWAP multiple calculations",
		testVWAPFunc(
			[]candle.Candle{
				{
					Volume:      89329,
					HighMicros:  utils.DollarsToMicros(127.36),
					LowMicros:   utils.DollarsToMicros(126.99),
					CloseMicros: utils.DollarsToMicros(127.28),
				},
				{
					Volume:      16137,
					HighMicros:  utils.DollarsToMicros(127.31),
					LowMicros:   utils.DollarsToMicros(127.10),
					CloseMicros: utils.DollarsToMicros(127.11),
				},
				{
					Volume:      23945,
					HighMicros:  utils.DollarsToMicros(127.21),
					LowMicros:   utils.DollarsToMicros(127.11),
					CloseMicros: utils.DollarsToMicros(127.15),
				},
				{
					Volume:      20679,
					HighMicros:  utils.DollarsToMicros(127.15),
					LowMicros:   utils.DollarsToMicros(126.93),
					CloseMicros: utils.DollarsToMicros(127.04),
				},
				{
					Volume:      27252,
					HighMicros:  utils.DollarsToMicros(127.08),
					LowMicros:   utils.DollarsToMicros(126.98),
					CloseMicros: utils.DollarsToMicros(126.98),
				},
				{
					Volume:      20915,
					HighMicros:  utils.DollarsToMicros(127.19),
					LowMicros:   utils.DollarsToMicros(126.99),
					CloseMicros: utils.DollarsToMicros(127.07),
				},
			},
			[]utils.MicroDollar{
				// (127.36 + 126.99 + 127.28) / 3 = 11
				utils.NewMicroDollar(utils.DollarsToMicros(127.21)),
				utils.NewMicroDollar(utils.DollarsToMicros(127.204389)),
				utils.NewMicroDollar(utils.DollarsToMicros(127.195559)),
				utils.NewMicroDollar(utils.DollarsToMicros(127.174126)),
				utils.NewMicroDollar(utils.DollarsToMicros(127.149417)),
				utils.NewMicroDollar(utils.DollarsToMicros(127.142446)),
			},
		),
	)
}

func testVWAPFunc(candles []candle.Candle, expected []utils.MicroDollar) func(*testing.T) {
	return func(t *testing.T) {
		actual := VWAP(candles)
		if !utils.EqMicroDollarSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

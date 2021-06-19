package analyze

import (
	"testing"

	"github.com/t73liu/tradingbot/lib/utils"
)

func TestEMA(t *testing.T) {
	t.Run(
		"EMA not enough elements for any calculation",
		testEMAFunc(
			[]int64{10, 10, 10},
			5,
			make([]utils.MicroDollar, 3),
		),
	)
	t.Run(
		"EMA enough elements for one calculation",
		testEMAFunc(
			[]int64{10, 10, 10, 10, 15},
			5,
			// Same as SMA for initial calc: (10 + 10 + 10 + 10 + 15) / 5 = 11
			[]utils.MicroDollar{{}, {}, {}, {}, utils.NewMicroDollar(11)},
		),
	)
	t.Run(
		"EMA enough elements for multiple calculations",
		testEMAFunc(
			[]int64{
				utils.DollarsToMicros(14),
				utils.DollarsToMicros(13),
				utils.DollarsToMicros(14),
				utils.DollarsToMicros(13),
				utils.DollarsToMicros(12),
				utils.DollarsToMicros(12),
				utils.DollarsToMicros(11),
			},
			5,
			[]utils.MicroDollar{
				{},
				{},
				{},
				{},
				// (14 + 13 + 14 + 13 + 12) / 5 = 13.2
				utils.NewMicroDollar(13200000),
				// (12 - 13.2) * (2 / (5 + 1)) + 13.2 = 12.8
				utils.NewMicroDollar(12800000),
				// (11 - 12.8) * (2 / (5 + 1)) + 12.8 = 12.2
				utils.NewMicroDollar(12200000),
			},
		),
	)
}

func testEMAFunc(closingPrices []int64, interval int, expected []utils.MicroDollar) func(*testing.T) {
	return func(t *testing.T) {
		actual := EMA(closingPrices, interval)
		if !utils.EqMicroDollarSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

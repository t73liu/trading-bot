package analyze

import (
	"testing"
	"tradingbot/lib/utils"
)

func TestSMA(t *testing.T) {
	t.Run(
		"SMA not enough elements for any calculation",
		testSMAFunc(
			[]int64{10, 10, 10},
			5,
			make([]utils.MicroDollar, 3),
		),
	)
	t.Run(
		"SMA enough elements for one calculation",
		testSMAFunc(
			[]int64{10, 10, 10, 10, 15},
			5,
			// (10 + 10 + 10 + 10 + 15) / 5 = 11
			[]utils.MicroDollar{{}, {}, {}, {}, utils.NewMicroDollar(11)},
		),
	)
	t.Run(
		"SMA enough elements for multiple calculations",
		testSMAFunc(
			[]int64{
				utils.DollarsToMicros(13),
				utils.DollarsToMicros(17),
				utils.DollarsToMicros(14),
				utils.DollarsToMicros(16),
				utils.DollarsToMicros(15),
				utils.DollarsToMicros(20),
				utils.DollarsToMicros(123),
			},
			5,
			[]utils.MicroDollar{
				{},
				{},
				{},
				{},
				// (13 + 17 + 14 + 16 + 15) / 5 = 15
				utils.NewMicroDollar(15000000),
				// (17 + 14 + 16 + 15 + 20) / 5 = 16.4
				utils.NewMicroDollar(16400000),
				// (14 + 16 + 15 + 20 + 123) / 5 = 37.6
				utils.NewMicroDollar(37600000),
			},
		),
	)
}

func testSMAFunc(closingPrices []int64, interval int, expected []utils.MicroDollar) func(*testing.T) {
	return func(t *testing.T) {
		actual := SMA(closingPrices, interval)
		if !utils.EqMicroDollarSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

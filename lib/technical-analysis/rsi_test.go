package analyze

import (
	"testing"

	"github.com/t73liu/tradingbot/lib/utils"
)

func TestRSI(t *testing.T) {
	t.Run(
		"RSI not enough elements for any calculation",
		testRSIFunc(
			[]int64{10, 10, 10},
			5,
			make([]utils.NullFloat64, 3),
		),
	)
	t.Run(
		"RSI not enough elements for any calculation (number of periods = interval)",
		testRSIFunc(
			[]int64{10, 10, 10, 10, 15},
			5,
			make([]utils.NullFloat64, 5),
		),
	)
	t.Run(
		"RSI enough elements for one calculation (number of periods = interval + 1)",
		testRSIFunc(
			[]int64{
				utils.DollarsToMicros(283.46),
				utils.DollarsToMicros(280.69),
				utils.DollarsToMicros(285.48),
				utils.DollarsToMicros(294.08),
				utils.DollarsToMicros(293.90),
				utils.DollarsToMicros(299.92),
				utils.DollarsToMicros(301.15),
				utils.DollarsToMicros(284.45),
				utils.DollarsToMicros(294.09),
				utils.DollarsToMicros(302.77),
				utils.DollarsToMicros(301.97),
				utils.DollarsToMicros(306.85),
				utils.DollarsToMicros(305.02),
				utils.DollarsToMicros(301.06),
				utils.DollarsToMicros(291.97),
			},
			14,
			[]utils.NullFloat64{
				{}, {}, {}, {}, {},
				{}, {}, {}, {}, {},
				{}, {}, {}, {},
				utils.NewNullFloat64(55.37),
			},
		),
	)
	t.Run(
		"RSI only gains",
		testRSIFunc(
			[]int64{
				utils.DollarsToMicros(13),
				utils.DollarsToMicros(14),
				utils.DollarsToMicros(15),
				utils.DollarsToMicros(15),
				utils.DollarsToMicros(15),
				utils.DollarsToMicros(15),
			},
			5,
			[]utils.NullFloat64{{}, {}, {}, {}, {}, utils.NewNullFloat64(100)},
		),
	)
	t.Run(
		"RSI only losses",
		testRSIFunc(
			[]int64{
				utils.DollarsToMicros(13),
				utils.DollarsToMicros(12),
				utils.DollarsToMicros(12),
				utils.DollarsToMicros(12),
				utils.DollarsToMicros(11),
				utils.DollarsToMicros(1),
			},
			5,
			[]utils.NullFloat64{{}, {}, {}, {}, {}, utils.NewNullFloat64(0)},
		),
	)
	t.Run(
		"RSI enough elements for multiple calculations",
		testRSIFunc(
			[]int64{
				utils.DollarsToMicros(283.46),
				utils.DollarsToMicros(280.69),
				utils.DollarsToMicros(285.48),
				utils.DollarsToMicros(294.08),
				utils.DollarsToMicros(293.90),
				utils.DollarsToMicros(299.92),
				utils.DollarsToMicros(301.15),
				utils.DollarsToMicros(284.45),
				utils.DollarsToMicros(294.09),
				utils.DollarsToMicros(302.77),
				utils.DollarsToMicros(301.97),
				utils.DollarsToMicros(306.85),
				utils.DollarsToMicros(305.02),
				utils.DollarsToMicros(301.06),
				utils.DollarsToMicros(291.97),
				utils.DollarsToMicros(284.18),
				utils.DollarsToMicros(286.48),
				utils.DollarsToMicros(284.54),
				utils.DollarsToMicros(276.82),
				utils.DollarsToMicros(284.49),
			},
			14,
			[]utils.NullFloat64{
				{}, {}, {}, {}, {},
				{}, {}, {}, {}, {},
				{}, {}, {}, {},
				utils.NewNullFloat64(55.37),
				utils.NewNullFloat64(50.07),
				utils.NewNullFloat64(51.55),
				utils.NewNullFloat64(50.20),
				utils.NewNullFloat64(45.14),
				utils.NewNullFloat64(50.48),
			},
		),
	)
}

func testRSIFunc(closingPrices []int64, interval int, expected []utils.NullFloat64) func(*testing.T) {
	return func(t *testing.T) {
		actual := RSI(closingPrices, interval)
		if !utils.EqNullFloat64Slice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

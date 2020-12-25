package analyze

import (
	"testing"
	"tradingbot/lib/utils"
)

func TestSmoothedRateOfChange(t *testing.T) {
	t.Run(
		"Smoothed Rate of Change not enough elements for any calculation",
		testSmoothedRateOfChangeFunc(
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
			},
			5,
			5,
			make([]utils.NullFloat64, 9),
		),
	)
	t.Run(
		"Smoothed Rate of Change enough elements for one calculation (number of periods = EMA interval + RoC interval)",
		testSmoothedRateOfChangeFunc(
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
			},
			5,
			5,
			[]utils.NullFloat64{
				{}, {}, {}, {}, {},
				{}, {}, {}, {},
				utils.NewNullFloat64(2.87),
			},
		),
	)
	t.Run(
		"Smoothed Rate of Change enough elements for multiple calculations",
		testSmoothedRateOfChangeFunc(
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
			},
			5,
			5,
			[]utils.NullFloat64{
				{}, {}, {}, {}, {},
				{}, {}, {}, {},
				utils.NewNullFloat64(2.87),
				utils.NewNullFloat64(2.12),
				utils.NewNullFloat64(2.04),
			},
		),
	)
}

func testSmoothedRateOfChangeFunc(closingPrices []int64, averageInterval, rateOfChangeInterval int, expected []utils.NullFloat64) func(*testing.T) {
	return func(t *testing.T) {
		actual := SmoothedRateOfChange(closingPrices, averageInterval, rateOfChangeInterval)
		if !utils.EqNullFloat64Slice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

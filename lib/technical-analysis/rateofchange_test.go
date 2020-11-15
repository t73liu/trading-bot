package analyze

import (
	"testing"
	"tradingbot/lib/candle"
)

func TestSmoothedRateOfChange(t *testing.T) {
	t.Run(
		"Smoothed Rate of Change not enough elements for any calculation",
		testSmoothedRateOfChangeFunc(
			[]int64{
				candle.DollarsToMicros(283.46),
				candle.DollarsToMicros(280.69),
				candle.DollarsToMicros(285.48),
				candle.DollarsToMicros(294.08),
				candle.DollarsToMicros(293.90),
				candle.DollarsToMicros(299.92),
				candle.DollarsToMicros(301.15),
				candle.DollarsToMicros(284.45),
				candle.DollarsToMicros(294.09),
			},
			5,
			5,
			make([]ValidFloat, 9),
		),
	)
	t.Run(
		"Smoothed Rate of Change enough elements for one calculation (number of periods = EMA interval + RoC interval)",
		testSmoothedRateOfChangeFunc(
			[]int64{
				candle.DollarsToMicros(283.46),
				candle.DollarsToMicros(280.69),
				candle.DollarsToMicros(285.48),
				candle.DollarsToMicros(294.08),
				candle.DollarsToMicros(293.90),
				candle.DollarsToMicros(299.92),
				candle.DollarsToMicros(301.15),
				candle.DollarsToMicros(284.45),
				candle.DollarsToMicros(294.09),
				candle.DollarsToMicros(302.77),
			},
			5,
			5,
			[]ValidFloat{
				{}, {}, {}, {}, {},
				{}, {}, {}, {},
				genValidFloat(2.87),
			},
		),
	)
	t.Run(
		"Smoothed Rate of Change enough elements for multiple calculations",
		testSmoothedRateOfChangeFunc(
			[]int64{
				candle.DollarsToMicros(283.46),
				candle.DollarsToMicros(280.69),
				candle.DollarsToMicros(285.48),
				candle.DollarsToMicros(294.08),
				candle.DollarsToMicros(293.90),
				candle.DollarsToMicros(299.92),
				candle.DollarsToMicros(301.15),
				candle.DollarsToMicros(284.45),
				candle.DollarsToMicros(294.09),
				candle.DollarsToMicros(302.77),
				candle.DollarsToMicros(301.97),
				candle.DollarsToMicros(306.85),
			},
			5,
			5,
			[]ValidFloat{
				{}, {}, {}, {}, {},
				{}, {}, {}, {},
				genValidFloat(2.87),
				genValidFloat(2.12),
				genValidFloat(2.04),
			},
		),
	)
}

func testSmoothedRateOfChangeFunc(closingPrices []int64, averageInterval, rateOfChangeInterval int, expected []ValidFloat) func(*testing.T) {
	return func(t *testing.T) {
		actual := SmoothedRateOfChange(closingPrices, averageInterval, rateOfChangeInterval)
		if !eqValidFloatSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

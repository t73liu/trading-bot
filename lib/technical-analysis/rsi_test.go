package analyze

import (
	"testing"
	"tradingbot/lib/candle"
)

func TestRSI(t *testing.T) {
	t.Run(
		"RSI not enough elements for any calculation",
		testRSIFunc(
			[]int64{10, 10, 10},
			5,
			make([]ValidFloat, 3),
		),
	)
	t.Run(
		"RSI not enough elements for any calculation (number of periods = interval)",
		testRSIFunc(
			[]int64{10, 10, 10, 10, 15},
			5,
			make([]ValidFloat, 5),
		),
	)
	t.Run(
		"RSI enough elements for one calculation (number of periods = interval + 1)",
		testRSIFunc(
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
				candle.DollarsToMicros(305.02),
				candle.DollarsToMicros(301.06),
				candle.DollarsToMicros(291.97),
			},
			14,
			[]ValidFloat{
				{}, {}, {}, {}, {},
				{}, {}, {}, {}, {},
				{}, {}, {}, {},
				genValidFloat(55.37),
			},
		),
	)
	t.Run(
		"RSI only gains",
		testRSIFunc(
			[]int64{
				candle.DollarsToMicros(13),
				candle.DollarsToMicros(14),
				candle.DollarsToMicros(15),
				candle.DollarsToMicros(15),
				candle.DollarsToMicros(15),
				candle.DollarsToMicros(15),
			},
			5,
			[]ValidFloat{{}, {}, {}, {}, {}, genValidFloat(100)},
		),
	)
	t.Run(
		"RSI only losses",
		testRSIFunc(
			[]int64{
				candle.DollarsToMicros(13),
				candle.DollarsToMicros(12),
				candle.DollarsToMicros(12),
				candle.DollarsToMicros(12),
				candle.DollarsToMicros(11),
				candle.DollarsToMicros(1),
			},
			5,
			[]ValidFloat{{}, {}, {}, {}, {}, genValidFloat(0)},
		),
	)
	t.Run(
		"RSI enough elements for multiple calculations",
		testRSIFunc(
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
				candle.DollarsToMicros(305.02),
				candle.DollarsToMicros(301.06),
				candle.DollarsToMicros(291.97),
				candle.DollarsToMicros(284.18),
				candle.DollarsToMicros(286.48),
				candle.DollarsToMicros(284.54),
				candle.DollarsToMicros(276.82),
				candle.DollarsToMicros(284.49),
			},
			14,
			[]ValidFloat{
				{}, {}, {}, {}, {},
				{}, {}, {}, {}, {},
				{}, {}, {}, {},
				genValidFloat(55.37),
				genValidFloat(50.07),
				genValidFloat(51.55),
				genValidFloat(50.20),
				genValidFloat(45.14),
				genValidFloat(50.48),
			},
		),
	)
}

func testRSIFunc(closingPrices []int64, interval int, expected []ValidFloat) func(*testing.T) {
	return func(t *testing.T) {
		actual := RSI(closingPrices, interval)
		if !eqValidFloatSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

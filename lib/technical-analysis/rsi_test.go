package analyze

import "testing"

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
				DollarsToMicros(283.46),
				DollarsToMicros(280.69),
				DollarsToMicros(285.48),
				DollarsToMicros(294.08),
				DollarsToMicros(293.90),
				DollarsToMicros(299.92),
				DollarsToMicros(301.15),
				DollarsToMicros(284.45),
				DollarsToMicros(294.09),
				DollarsToMicros(302.77),
				DollarsToMicros(301.97),
				DollarsToMicros(306.85),
				DollarsToMicros(305.02),
				DollarsToMicros(301.06),
				DollarsToMicros(291.97),
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
				DollarsToMicros(13),
				DollarsToMicros(14),
				DollarsToMicros(15),
				DollarsToMicros(15),
				DollarsToMicros(15),
				DollarsToMicros(15),
			},
			5,
			[]ValidFloat{{}, {}, {}, {}, {}, genValidFloat(100)},
		),
	)
	t.Run(
		"RSI only losses",
		testRSIFunc(
			[]int64{
				DollarsToMicros(13),
				DollarsToMicros(12),
				DollarsToMicros(12),
				DollarsToMicros(12),
				DollarsToMicros(11),
				DollarsToMicros(1),
			},
			5,
			[]ValidFloat{{}, {}, {}, {}, {}, genValidFloat(0)},
		),
	)
	t.Run(
		"RSI enough elements for multiple calculations",
		testRSIFunc(
			[]int64{
				DollarsToMicros(283.46),
				DollarsToMicros(280.69),
				DollarsToMicros(285.48),
				DollarsToMicros(294.08),
				DollarsToMicros(293.90),
				DollarsToMicros(299.92),
				DollarsToMicros(301.15),
				DollarsToMicros(284.45),
				DollarsToMicros(294.09),
				DollarsToMicros(302.77),
				DollarsToMicros(301.97),
				DollarsToMicros(306.85),
				DollarsToMicros(305.02),
				DollarsToMicros(301.06),
				DollarsToMicros(291.97),
				DollarsToMicros(284.18),
				DollarsToMicros(286.48),
				DollarsToMicros(284.54),
				DollarsToMicros(276.82),
				DollarsToMicros(284.49),
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

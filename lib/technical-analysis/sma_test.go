package analyze

import "testing"

func TestSMA(t *testing.T) {
	t.Run(
		"SMA not enough elements for any calculation",
		testSMAFunc(
			[]int64{10, 10, 10},
			5,
			make([]ValidMicro, 3),
		),
	)
	t.Run(
		"SMA enough elements for one calculation",
		testSMAFunc(
			[]int64{10, 10, 10, 10, 15},
			5,
			// (10 + 10 + 10 + 10 + 15) / 5 = 11
			[]ValidMicro{{}, {}, {}, {}, genValidMicro(11)},
		),
	)
	t.Run(
		"SMA enough elements for multiple calculations",
		testSMAFunc(
			[]int64{
				DollarsToMicros(13),
				DollarsToMicros(17),
				DollarsToMicros(14),
				DollarsToMicros(16),
				DollarsToMicros(15),
				DollarsToMicros(20),
				DollarsToMicros(123),
			},
			5,
			[]ValidMicro{
				{},
				{},
				{},
				{},
				// (13 + 17 + 14 + 16 + 15) / 5 = 15
				genValidMicro(15000000),
				// (17 + 14 + 16 + 15 + 20) / 5 = 16.4
				genValidMicro(16400000),
				// (14 + 16 + 15 + 20 + 123) / 5 = 37.6
				genValidMicro(37600000),
			},
		),
	)
}

func testSMAFunc(closingPrices []int64, interval int, expected []ValidMicro) func(*testing.T) {
	return func(t *testing.T) {
		actual := SMA(closingPrices, interval)
		if !eqValidCalcSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

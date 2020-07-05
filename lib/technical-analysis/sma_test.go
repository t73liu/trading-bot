package analyze

import "testing"

func TestSMA(t *testing.T) {
	t.Run(
		"SMA not enough elements for any calculation",
		testSMAFunc(
			[]Candle{{Close: 10}, {Close: 10}, {Close: 10}},
			5,
			make([]ValidCalc, 3),
		),
	)
	t.Run(
		"SMA enough elements for one calculation",
		testSMAFunc(
			[]Candle{{Close: 10}, {Close: 10}, {Close: 10}, {Close: 10}, {Close: 15}},
			5,
			[]ValidCalc{{}, {}, {}, {}, genValidCalc(11)},
		),
	)
	t.Run(
		"SMA enough elements for multiple calculations",
		testSMAFunc(
			[]Candle{
				{Close: dollarsToMicros(13)},
				{Close: dollarsToMicros(17)},
				{Close: dollarsToMicros(14)},
				{Close: dollarsToMicros(16)},
				{Close: dollarsToMicros(15)},
				{Close: dollarsToMicros(20)},
				{Close: dollarsToMicros(123)},
			},
			5,
			[]ValidCalc{
				{},
				{},
				{},
				{},
				// (13 + 17 + 14 + 16 + 15) / 5 = 15
				genValidCalc(15000000),
				// (17 + 14 + 16 + 15 + 20) / 5 = 16.4
				genValidCalc(16400000),
				// (14 + 16 + 15 + 20 + 123) / 5 = 37.6
				genValidCalc(37600000),
			},
		),
	)
}

func testSMAFunc(candles []Candle, interval int, expected []ValidCalc) func(*testing.T) {
	return func(t *testing.T) {
		actual := SMA(candles, interval)
		if !eqValidCalcSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

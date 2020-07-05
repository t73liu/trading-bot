package analyze

import "testing"

func TestEMA(t *testing.T) {
	t.Run(
		"EMA not enough elements for any calculation",
		testEMAFunc(
			[]Candle{{Close: 10}, {Close: 10}, {Close: 10}},
			5,
			make([]ValidCalc, 3),
		),
	)
	t.Run(
		"EMA enough elements for one calculation",
		testEMAFunc(
			[]Candle{{Close: 10}, {Close: 10}, {Close: 10}, {Close: 10}, {Close: 15}},
			5,
			[]ValidCalc{{}, {}, {}, {}, genValidCalc(11)},
		),
	)
	t.Run(
		"EMA enough elements for multiple calculations",
		testEMAFunc(
			[]Candle{
				{Close: dollarsToMicros(14)},
				{Close: dollarsToMicros(13)},
				{Close: dollarsToMicros(14)},
				{Close: dollarsToMicros(13)},
				{Close: dollarsToMicros(12)},
				{Close: dollarsToMicros(12)},
				{Close: dollarsToMicros(11)},
			},
			5,
			[]ValidCalc{
				{},
				{},
				{},
				{},
				// (14 + 13 + 14 + 13 + 12) / 5 = 13.2
				genValidCalc(13200000),
				// (12 - 13.2) * (2 / (5 + 1)) + 13.2 = 12.8
				genValidCalc(12800000),
				// (11 - 12.8) * (2 / (5 + 1)) + 12.8 = 12.2
				genValidCalc(12200000),
			},
		),
	)
}

func testEMAFunc(candles []Candle, interval int, expected []ValidCalc) func(*testing.T) {
	return func(t *testing.T) {
		actual := EMA(candles, interval)
		if !eqValidCalcSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

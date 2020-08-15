package analyze

import (
	"testing"
	"tradingbot/lib/candle"
)

func TestBollingerBands(t *testing.T) {
	t.Run(
		"Bollinger Bands not enough elements for any calculation",
		testBollingerBandsFunc(
			[]int64{50, 4000, 357},
			make([]ValidMicroRange, 3),
		),
	)
	t.Run(
		"Bollinger Bands enough elements for one calculation",
		testBollingerBandsFunc(
			[]int64{
				candle.DollarsToMicros(86.16),
				candle.DollarsToMicros(89.09),
				candle.DollarsToMicros(88.78),
				candle.DollarsToMicros(90.32),
				candle.DollarsToMicros(89.07),
				candle.DollarsToMicros(91.15),
				candle.DollarsToMicros(89.44),
				candle.DollarsToMicros(89.18),
				candle.DollarsToMicros(86.93),
				candle.DollarsToMicros(87.68),
				candle.DollarsToMicros(86.96),
				candle.DollarsToMicros(89.43),
				candle.DollarsToMicros(89.32),
				candle.DollarsToMicros(88.72),
				candle.DollarsToMicros(87.45),
				candle.DollarsToMicros(87.26),
				candle.DollarsToMicros(89.50),
				candle.DollarsToMicros(87.90),
				candle.DollarsToMicros(89.13),
				candle.DollarsToMicros(90.70),
			},
			append(
				make([]ValidMicroRange, 19, 20),
				ValidMicroRange{
					Valid: true,
					High:  91291910,
					Mid:   88708500,
					Low:   86125090,
				},
			),
		),
	)
	t.Run(
		"Bollinger Bands enough elements for multiple calculations",
		testBollingerBandsFunc(
			[]int64{
				candle.DollarsToMicros(86.16),
				candle.DollarsToMicros(89.09),
				candle.DollarsToMicros(88.78),
				candle.DollarsToMicros(90.32),
				candle.DollarsToMicros(89.07),
				candle.DollarsToMicros(91.15),
				candle.DollarsToMicros(89.44),
				candle.DollarsToMicros(89.18),
				candle.DollarsToMicros(86.93),
				candle.DollarsToMicros(87.68),
				candle.DollarsToMicros(86.96),
				candle.DollarsToMicros(89.43),
				candle.DollarsToMicros(89.32),
				candle.DollarsToMicros(88.72),
				candle.DollarsToMicros(87.45),
				candle.DollarsToMicros(87.26),
				candle.DollarsToMicros(89.50),
				candle.DollarsToMicros(87.90),
				candle.DollarsToMicros(89.13),
				candle.DollarsToMicros(90.70),
				candle.DollarsToMicros(92.9),
				candle.DollarsToMicros(92.98),
			},
			append(
				make([]ValidMicroRange, 19, 20),
				ValidMicroRange{
					Valid: true,
					High:  91291910,
					Mid:   88708500,
					Low:   86125090,
				},
				ValidMicroRange{
					Valid: true,
					High:  91949720,
					Mid:   89045500,
					Low:   86141280,
				},
				ValidMicroRange{
					Valid: true,
					High:  92613252,
					Mid:   89240000,
					Low:   85866748,
				},
			),
		),
	)
}

func testBollingerBandsFunc(closingPrices []int64, expected []ValidMicroRange) func(*testing.T) {
	return func(t *testing.T) {
		actual := BollingerBands(closingPrices, 20)
		if !eqValidMicroRangeSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

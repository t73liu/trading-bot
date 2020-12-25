package analyze

import (
	"testing"
	"tradingbot/lib/utils"
)

func TestBollingerBands(t *testing.T) {
	t.Run(
		"Bollinger Bands not enough elements for any calculation",
		testBollingerBandsFunc(
			[]int64{50, 4000, 357},
			make([]MicroDollarRange, 3),
		),
	)
	t.Run(
		"Bollinger Bands enough elements for one calculation",
		testBollingerBandsFunc(
			[]int64{
				utils.DollarsToMicros(86.16),
				utils.DollarsToMicros(89.09),
				utils.DollarsToMicros(88.78),
				utils.DollarsToMicros(90.32),
				utils.DollarsToMicros(89.07),
				utils.DollarsToMicros(91.15),
				utils.DollarsToMicros(89.44),
				utils.DollarsToMicros(89.18),
				utils.DollarsToMicros(86.93),
				utils.DollarsToMicros(87.68),
				utils.DollarsToMicros(86.96),
				utils.DollarsToMicros(89.43),
				utils.DollarsToMicros(89.32),
				utils.DollarsToMicros(88.72),
				utils.DollarsToMicros(87.45),
				utils.DollarsToMicros(87.26),
				utils.DollarsToMicros(89.50),
				utils.DollarsToMicros(87.90),
				utils.DollarsToMicros(89.13),
				utils.DollarsToMicros(90.70),
			},
			append(
				make([]MicroDollarRange, 19, 20),
				MicroDollarRange{
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
				utils.DollarsToMicros(86.16),
				utils.DollarsToMicros(89.09),
				utils.DollarsToMicros(88.78),
				utils.DollarsToMicros(90.32),
				utils.DollarsToMicros(89.07),
				utils.DollarsToMicros(91.15),
				utils.DollarsToMicros(89.44),
				utils.DollarsToMicros(89.18),
				utils.DollarsToMicros(86.93),
				utils.DollarsToMicros(87.68),
				utils.DollarsToMicros(86.96),
				utils.DollarsToMicros(89.43),
				utils.DollarsToMicros(89.32),
				utils.DollarsToMicros(88.72),
				utils.DollarsToMicros(87.45),
				utils.DollarsToMicros(87.26),
				utils.DollarsToMicros(89.50),
				utils.DollarsToMicros(87.90),
				utils.DollarsToMicros(89.13),
				utils.DollarsToMicros(90.70),
				utils.DollarsToMicros(92.9),
				utils.DollarsToMicros(92.98),
			},
			append(
				make([]MicroDollarRange, 19, 20),
				MicroDollarRange{
					Valid: true,
					High:  91291910,
					Mid:   88708500,
					Low:   86125090,
				},
				MicroDollarRange{
					Valid: true,
					High:  91949720,
					Mid:   89045500,
					Low:   86141280,
				},
				MicroDollarRange{
					Valid: true,
					High:  92613252,
					Mid:   89240000,
					Low:   85866748,
				},
			),
		),
	)
}

func testBollingerBandsFunc(closingPrices []int64, expected []MicroDollarRange) func(*testing.T) {
	return func(t *testing.T) {
		actual := BollingerBands(closingPrices, 20)
		if !eqMicroDollarRangeSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

package analyze

import (
	"testing"
	"tradingbot/lib/candle"
	"tradingbot/lib/utils"
)

func TestATR(t *testing.T) {
	t.Run(
		"ATR not enough elements for calculation",
		testATRFunc(
			[]candle.Candle{
				{
					HighMicros:  utils.DollarsToMicros(48.7),
					LowMicros:   utils.DollarsToMicros(47.79),
					CloseMicros: utils.DollarsToMicros(48.16),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.72),
					LowMicros:   utils.DollarsToMicros(48.14),
					CloseMicros: utils.DollarsToMicros(48.61),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.9),
					LowMicros:   utils.DollarsToMicros(48.39),
					CloseMicros: utils.DollarsToMicros(48.75),
				},
			},
			make([]utils.MicroDollar, 3),
		),
	)
	t.Run(
		"ATR enough elements for one calculation",
		testATRFunc(
			[]candle.Candle{
				{
					HighMicros:  utils.DollarsToMicros(48.7),
					LowMicros:   utils.DollarsToMicros(47.79),
					CloseMicros: utils.DollarsToMicros(48.16),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.72),
					LowMicros:   utils.DollarsToMicros(48.14),
					CloseMicros: utils.DollarsToMicros(48.61),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.9),
					LowMicros:   utils.DollarsToMicros(48.39),
					CloseMicros: utils.DollarsToMicros(48.75),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.87),
					LowMicros:   utils.DollarsToMicros(48.37),
					CloseMicros: utils.DollarsToMicros(48.63),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.82),
					LowMicros:   utils.DollarsToMicros(48.24),
					CloseMicros: utils.DollarsToMicros(48.74),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.05),
					LowMicros:   utils.DollarsToMicros(48.64),
					CloseMicros: utils.DollarsToMicros(49.03),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.20),
					LowMicros:   utils.DollarsToMicros(48.94),
					CloseMicros: utils.DollarsToMicros(49.07),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.35),
					LowMicros:   utils.DollarsToMicros(48.86),
					CloseMicros: utils.DollarsToMicros(49.32),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.92),
					LowMicros:   utils.DollarsToMicros(49.5),
					CloseMicros: utils.DollarsToMicros(49.91),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.19),
					LowMicros:   utils.DollarsToMicros(49.87),
					CloseMicros: utils.DollarsToMicros(50.13),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.12),
					LowMicros:   utils.DollarsToMicros(49.20),
					CloseMicros: utils.DollarsToMicros(49.53),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.66),
					LowMicros:   utils.DollarsToMicros(48.9),
					CloseMicros: utils.DollarsToMicros(49.5),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.88),
					LowMicros:   utils.DollarsToMicros(49.43),
					CloseMicros: utils.DollarsToMicros(49.75),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.19),
					LowMicros:   utils.DollarsToMicros(49.73),
					CloseMicros: utils.DollarsToMicros(50.03),
				},
			},
			append(
				make([]utils.MicroDollar, 13, 14),
				utils.NewMicroDollar(554285),
			),
		),
	)
	t.Run(
		"ATR enough elements for multiple calculations",
		testATRFunc(
			[]candle.Candle{
				{
					HighMicros:  utils.DollarsToMicros(48.7),
					LowMicros:   utils.DollarsToMicros(47.79),
					CloseMicros: utils.DollarsToMicros(48.16),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.72),
					LowMicros:   utils.DollarsToMicros(48.14),
					CloseMicros: utils.DollarsToMicros(48.61),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.9),
					LowMicros:   utils.DollarsToMicros(48.39),
					CloseMicros: utils.DollarsToMicros(48.75),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.87),
					LowMicros:   utils.DollarsToMicros(48.37),
					CloseMicros: utils.DollarsToMicros(48.63),
				},
				{
					HighMicros:  utils.DollarsToMicros(48.82),
					LowMicros:   utils.DollarsToMicros(48.24),
					CloseMicros: utils.DollarsToMicros(48.74),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.05),
					LowMicros:   utils.DollarsToMicros(48.64),
					CloseMicros: utils.DollarsToMicros(49.03),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.20),
					LowMicros:   utils.DollarsToMicros(48.94),
					CloseMicros: utils.DollarsToMicros(49.07),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.35),
					LowMicros:   utils.DollarsToMicros(48.86),
					CloseMicros: utils.DollarsToMicros(49.32),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.92),
					LowMicros:   utils.DollarsToMicros(49.5),
					CloseMicros: utils.DollarsToMicros(49.91),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.19),
					LowMicros:   utils.DollarsToMicros(49.87),
					CloseMicros: utils.DollarsToMicros(50.13),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.12),
					LowMicros:   utils.DollarsToMicros(49.20),
					CloseMicros: utils.DollarsToMicros(49.53),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.66),
					LowMicros:   utils.DollarsToMicros(48.9),
					CloseMicros: utils.DollarsToMicros(49.5),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.88),
					LowMicros:   utils.DollarsToMicros(49.43),
					CloseMicros: utils.DollarsToMicros(49.75),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.19),
					LowMicros:   utils.DollarsToMicros(49.73),
					CloseMicros: utils.DollarsToMicros(50.03),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.36),
					LowMicros:   utils.DollarsToMicros(49.26),
					CloseMicros: utils.DollarsToMicros(50.31),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.57),
					LowMicros:   utils.DollarsToMicros(50.09),
					CloseMicros: utils.DollarsToMicros(50.52),
				},
			},
			append(
				make([]utils.MicroDollar, 13, 14),
				utils.NewMicroDollar(554285),
				utils.NewMicroDollar(593264),
				utils.NewMicroDollar(585173),
			),
		),
	)
}

func testATRFunc(candles []candle.Candle, expected []utils.MicroDollar) func(*testing.T) {
	return func(t *testing.T) {
		actual := ATR(candles, 14)
		if !utils.EqMicroDollarSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

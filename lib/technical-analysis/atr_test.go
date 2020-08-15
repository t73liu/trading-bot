package analyze

import (
	"testing"
	"tradingbot/lib/candle"
)

func TestATR(t *testing.T) {
	t.Run(
		"ATR not enough elements for calculation",
		testATRFunc(
			[]candle.Candle{
				{
					HighMicros:  candle.DollarsToMicros(48.7),
					LowMicros:   candle.DollarsToMicros(47.79),
					CloseMicros: candle.DollarsToMicros(48.16),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.72),
					LowMicros:   candle.DollarsToMicros(48.14),
					CloseMicros: candle.DollarsToMicros(48.61),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.9),
					LowMicros:   candle.DollarsToMicros(48.39),
					CloseMicros: candle.DollarsToMicros(48.75),
				},
			},
			make([]ValidMicro, 3),
		),
	)
	t.Run(
		"ATR enough elements for one calculation",
		testATRFunc(
			[]candle.Candle{
				{
					HighMicros:  candle.DollarsToMicros(48.7),
					LowMicros:   candle.DollarsToMicros(47.79),
					CloseMicros: candle.DollarsToMicros(48.16),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.72),
					LowMicros:   candle.DollarsToMicros(48.14),
					CloseMicros: candle.DollarsToMicros(48.61),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.9),
					LowMicros:   candle.DollarsToMicros(48.39),
					CloseMicros: candle.DollarsToMicros(48.75),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.87),
					LowMicros:   candle.DollarsToMicros(48.37),
					CloseMicros: candle.DollarsToMicros(48.63),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.82),
					LowMicros:   candle.DollarsToMicros(48.24),
					CloseMicros: candle.DollarsToMicros(48.74),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.05),
					LowMicros:   candle.DollarsToMicros(48.64),
					CloseMicros: candle.DollarsToMicros(49.03),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.20),
					LowMicros:   candle.DollarsToMicros(48.94),
					CloseMicros: candle.DollarsToMicros(49.07),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.35),
					LowMicros:   candle.DollarsToMicros(48.86),
					CloseMicros: candle.DollarsToMicros(49.32),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.92),
					LowMicros:   candle.DollarsToMicros(49.5),
					CloseMicros: candle.DollarsToMicros(49.91),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.19),
					LowMicros:   candle.DollarsToMicros(49.87),
					CloseMicros: candle.DollarsToMicros(50.13),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.12),
					LowMicros:   candle.DollarsToMicros(49.20),
					CloseMicros: candle.DollarsToMicros(49.53),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.66),
					LowMicros:   candle.DollarsToMicros(48.9),
					CloseMicros: candle.DollarsToMicros(49.5),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.88),
					LowMicros:   candle.DollarsToMicros(49.43),
					CloseMicros: candle.DollarsToMicros(49.75),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.19),
					LowMicros:   candle.DollarsToMicros(49.73),
					CloseMicros: candle.DollarsToMicros(50.03),
				},
			},
			append(
				make([]ValidMicro, 13, 14),
				genValidMicro(554285),
			),
		),
	)
	t.Run(
		"ATR enough elements for multiple calculations",
		testATRFunc(
			[]candle.Candle{
				{
					HighMicros:  candle.DollarsToMicros(48.7),
					LowMicros:   candle.DollarsToMicros(47.79),
					CloseMicros: candle.DollarsToMicros(48.16),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.72),
					LowMicros:   candle.DollarsToMicros(48.14),
					CloseMicros: candle.DollarsToMicros(48.61),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.9),
					LowMicros:   candle.DollarsToMicros(48.39),
					CloseMicros: candle.DollarsToMicros(48.75),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.87),
					LowMicros:   candle.DollarsToMicros(48.37),
					CloseMicros: candle.DollarsToMicros(48.63),
				},
				{
					HighMicros:  candle.DollarsToMicros(48.82),
					LowMicros:   candle.DollarsToMicros(48.24),
					CloseMicros: candle.DollarsToMicros(48.74),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.05),
					LowMicros:   candle.DollarsToMicros(48.64),
					CloseMicros: candle.DollarsToMicros(49.03),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.20),
					LowMicros:   candle.DollarsToMicros(48.94),
					CloseMicros: candle.DollarsToMicros(49.07),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.35),
					LowMicros:   candle.DollarsToMicros(48.86),
					CloseMicros: candle.DollarsToMicros(49.32),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.92),
					LowMicros:   candle.DollarsToMicros(49.5),
					CloseMicros: candle.DollarsToMicros(49.91),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.19),
					LowMicros:   candle.DollarsToMicros(49.87),
					CloseMicros: candle.DollarsToMicros(50.13),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.12),
					LowMicros:   candle.DollarsToMicros(49.20),
					CloseMicros: candle.DollarsToMicros(49.53),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.66),
					LowMicros:   candle.DollarsToMicros(48.9),
					CloseMicros: candle.DollarsToMicros(49.5),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.88),
					LowMicros:   candle.DollarsToMicros(49.43),
					CloseMicros: candle.DollarsToMicros(49.75),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.19),
					LowMicros:   candle.DollarsToMicros(49.73),
					CloseMicros: candle.DollarsToMicros(50.03),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.36),
					LowMicros:   candle.DollarsToMicros(49.26),
					CloseMicros: candle.DollarsToMicros(50.31),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.57),
					LowMicros:   candle.DollarsToMicros(50.09),
					CloseMicros: candle.DollarsToMicros(50.52),
				},
			},
			append(
				make([]ValidMicro, 13, 14),
				genValidMicro(554285),
				genValidMicro(593264),
				genValidMicro(585173),
			),
		),
	)
}

func testATRFunc(candles []candle.Candle, expected []ValidMicro) func(*testing.T) {
	return func(t *testing.T) {
		actual := ATR(candles, 14)
		if !eqValidMicroSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

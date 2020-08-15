package analyze

import (
	"testing"
	"tradingbot/lib/candle"
)

func TestKeltnerChannels(t *testing.T) {
	t.Run(
		"Keltner Channels not enough elements for calculation",
		testKeltnerChannelsFunc(
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
			make([]ValidMicroRange, 3),
		),
	)
	t.Run(
		"Keltner Channels enough elements for calculation",
		testKeltnerChannelsFunc(
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
				{
					HighMicros:  candle.DollarsToMicros(50.65),
					LowMicros:   candle.DollarsToMicros(50.3),
					CloseMicros: candle.DollarsToMicros(50.41),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.43),
					LowMicros:   candle.DollarsToMicros(49.21),
					CloseMicros: candle.DollarsToMicros(49.34),
				},
				{
					HighMicros:  candle.DollarsToMicros(49.63),
					LowMicros:   candle.DollarsToMicros(48.98),
					CloseMicros: candle.DollarsToMicros(49.37),
				},
				{
					HighMicros:  candle.DollarsToMicros(50.33),
					LowMicros:   candle.DollarsToMicros(49.61),
					CloseMicros: candle.DollarsToMicros(50.23),
				},
			},
			append(
				make([]ValidMicroRange, 19, 20),
				ValidMicroRange{
					Valid: true,
					High:  50806824,
					Mid:   49467000,
					Low:   48127176,
				},
			),
		),
	)
}

func testKeltnerChannelsFunc(candles []candle.Candle, expected []ValidMicroRange) func(*testing.T) {
	return func(t *testing.T) {
		actual := KeltnerChannels(candles, 20, 10)
		if !eqValidMicroRangeSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

package analyze

import (
	"testing"

	"github.com/t73liu/tradingbot/lib/candle"
	"github.com/t73liu/tradingbot/lib/utils"
)

func TestTTMSqueeze(t *testing.T) {
	t.Run(
		"TTM Squeeze not enough elements for any calculation",
		testTTMSqueezeFunc(
			make([]candle.Candle, 3),
			make([]utils.NullBool, 3),
		),
	)
	t.Run(
		"TTM Squeeze enough elements for sample calculation",
		testTTMSqueezeFunc(
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
				{
					HighMicros:  utils.DollarsToMicros(50.65),
					LowMicros:   utils.DollarsToMicros(50.3),
					CloseMicros: utils.DollarsToMicros(50.41),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.43),
					LowMicros:   utils.DollarsToMicros(49.21),
					CloseMicros: utils.DollarsToMicros(49.34),
				},
				{
					HighMicros:  utils.DollarsToMicros(49.63),
					LowMicros:   utils.DollarsToMicros(48.98),
					CloseMicros: utils.DollarsToMicros(49.37),
				},
				{
					HighMicros:  utils.DollarsToMicros(50.33),
					LowMicros:   utils.DollarsToMicros(49.61),
					CloseMicros: utils.DollarsToMicros(50.23),
				},
			},
			append(
				make([]utils.NullBool, 19, 20),
				utils.NewNullBool(true),
			),
		),
	)
}

func testTTMSqueezeFunc(candles []candle.Candle, expected []utils.NullBool) func(*testing.T) {
	return func(t *testing.T) {
		actual := TTMSqueeze(candles)
		if !utils.EqNullBoolSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

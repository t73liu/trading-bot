package analyze

import (
	"testing"
	"tradingbot/lib/candle"
)

func TestVWAP(t *testing.T) {
	t.Run(
		"VWAP empty candles returns nil",
		testVWAPFunc([]candle.Candle{}, nil),
	)
	t.Run(
		"VWAP first value is average of high, low, close",
		testVWAPFunc(
			[]candle.Candle{
				{
					Volume:      89329,
					HighMicros:  DollarsToMicros(127.36),
					LowMicros:   DollarsToMicros(126.99),
					CloseMicros: DollarsToMicros(127.28),
				},
			},
			// (127.36 + 126.99 + 127.28) / 3 = 11
			[]ValidMicro{genValidMicro(DollarsToMicros(127.21))},
		),
	)
	t.Run(
		"VWAP multiple calculations",
		testVWAPFunc(
			[]candle.Candle{
				{
					Volume:      89329,
					HighMicros:  DollarsToMicros(127.36),
					LowMicros:   DollarsToMicros(126.99),
					CloseMicros: DollarsToMicros(127.28),
				},
				{
					Volume:      16137,
					HighMicros:  DollarsToMicros(127.31),
					LowMicros:   DollarsToMicros(127.10),
					CloseMicros: DollarsToMicros(127.11),
				},
				{
					Volume:      23945,
					HighMicros:  DollarsToMicros(127.21),
					LowMicros:   DollarsToMicros(127.11),
					CloseMicros: DollarsToMicros(127.15),
				},
				{
					Volume:      20679,
					HighMicros:  DollarsToMicros(127.15),
					LowMicros:   DollarsToMicros(126.93),
					CloseMicros: DollarsToMicros(127.04),
				},
				{
					Volume:      27252,
					HighMicros:  DollarsToMicros(127.08),
					LowMicros:   DollarsToMicros(126.98),
					CloseMicros: DollarsToMicros(126.98),
				},
				{
					Volume:      20915,
					HighMicros:  DollarsToMicros(127.19),
					LowMicros:   DollarsToMicros(126.99),
					CloseMicros: DollarsToMicros(127.07),
				},
			},
			[]ValidMicro{
				// (127.36 + 126.99 + 127.28) / 3 = 11
				genValidMicro(DollarsToMicros(127.21)),
				genValidMicro(DollarsToMicros(127.204389)),
				genValidMicro(DollarsToMicros(127.195559)),
				genValidMicro(DollarsToMicros(127.174126)),
				genValidMicro(DollarsToMicros(127.149417)),
				genValidMicro(DollarsToMicros(127.142446)),
			},
		),
	)
}

func testVWAPFunc(candles []candle.Candle, expected []ValidMicro) func(*testing.T) {
	return func(t *testing.T) {
		actual := VWAP(candles)
		if !eqValidCalcSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

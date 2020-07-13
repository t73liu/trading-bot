package analyze

import (
	"testing"
)

func TestVWAP(t *testing.T) {
	t.Run(
		"VWAP empty candles returns nil",
		testVWAPFunc([]Candle{}, nil),
	)
	t.Run(
		"VWAP first value is average of high, low, close",
		testVWAPFunc(
			[]Candle{
				{
					Volume: 89329,
					High:   DollarsToMicros(127.36),
					Low:    DollarsToMicros(126.99),
					Close:  DollarsToMicros(127.28),
				},
			},
			// (127.36 + 126.99 + 127.28) / 3 = 11
			[]ValidMicro{genValidMicro(DollarsToMicros(127.21))},
		),
	)
	t.Run(
		"VWAP multiple calculations",
		testVWAPFunc(
			[]Candle{
				{
					Volume: 89329,
					High:   DollarsToMicros(127.36),
					Low:    DollarsToMicros(126.99),
					Close:  DollarsToMicros(127.28),
				},
				{
					Volume: 16137,
					High:   DollarsToMicros(127.31),
					Low:    DollarsToMicros(127.10),
					Close:  DollarsToMicros(127.11),
				},
				{
					Volume: 23945,
					High:   DollarsToMicros(127.21),
					Low:    DollarsToMicros(127.11),
					Close:  DollarsToMicros(127.15),
				},
				{
					Volume: 20679,
					High:   DollarsToMicros(127.15),
					Low:    DollarsToMicros(126.93),
					Close:  DollarsToMicros(127.04),
				},
				{
					Volume: 27252,
					High:   DollarsToMicros(127.08),
					Low:    DollarsToMicros(126.98),
					Close:  DollarsToMicros(126.98),
				},
				{
					Volume: 20915,
					High:   DollarsToMicros(127.19),
					Low:    DollarsToMicros(126.99),
					Close:  DollarsToMicros(127.07),
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

func testVWAPFunc(candles []Candle, expected []ValidMicro) func(*testing.T) {
	return func(t *testing.T) {
		actual := VWAP(candles)
		if !eqValidCalcSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

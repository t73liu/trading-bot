package analyze

import (
	"testing"
	"tradingbot/lib/candle"
)

func TestStandardMACD(t *testing.T) {
	t.Run(
		"Standard MACD not enough elements for any calculation",
		testStandardMACDFunc(
			[]int64{10, 10, 10, 10},
			make([]ValidMicro, 4),
		),
	)
	t.Run(
		"Standard MACD enough elements for one calculation",
		testStandardMACDFunc(
			[]int64{
				candle.DollarsToMicros(459.99),
				candle.DollarsToMicros(448.85),
				candle.DollarsToMicros(446.06),
				candle.DollarsToMicros(450.81),
				candle.DollarsToMicros(442.8),
				candle.DollarsToMicros(448.97),
				candle.DollarsToMicros(444.57),
				candle.DollarsToMicros(441.4),
				candle.DollarsToMicros(430.47),
				candle.DollarsToMicros(420.05),
				candle.DollarsToMicros(431.14),
				candle.DollarsToMicros(425.66),
				candle.DollarsToMicros(430.58),
				candle.DollarsToMicros(431.72),
				candle.DollarsToMicros(437.87),
				candle.DollarsToMicros(428.43),
				candle.DollarsToMicros(428.35),
				candle.DollarsToMicros(432.5),
				candle.DollarsToMicros(443.66),
				candle.DollarsToMicros(455.72),
				candle.DollarsToMicros(454.49),
				candle.DollarsToMicros(452.08),
				candle.DollarsToMicros(452.73),
				candle.DollarsToMicros(461.91),
				candle.DollarsToMicros(463.58),
				candle.DollarsToMicros(461.14),
				candle.DollarsToMicros(452.08),
				candle.DollarsToMicros(442.66),
				candle.DollarsToMicros(428.91),
				candle.DollarsToMicros(429.79),
				candle.DollarsToMicros(431.99),
				candle.DollarsToMicros(427.72),
				candle.DollarsToMicros(423.2),
				candle.DollarsToMicros(426.21),
			},
			// Same as SMA for initial calc: (10 + 10 + 10 + 10 + 15) / 5 = 11
			append(
				make([]ValidMicro, 33, 34),
				genValidMicro(-5108084),
			),
		),
	)
	t.Run(
		"Standard MACD enough elements for multiple calculations",
		testStandardMACDFunc(
			[]int64{
				candle.DollarsToMicros(459.99),
				candle.DollarsToMicros(448.85),
				candle.DollarsToMicros(446.06),
				candle.DollarsToMicros(450.81),
				candle.DollarsToMicros(442.8),
				candle.DollarsToMicros(448.97),
				candle.DollarsToMicros(444.57),
				candle.DollarsToMicros(441.4),
				candle.DollarsToMicros(430.47),
				candle.DollarsToMicros(420.05),
				candle.DollarsToMicros(431.14),
				candle.DollarsToMicros(425.66),
				candle.DollarsToMicros(430.58),
				candle.DollarsToMicros(431.72),
				candle.DollarsToMicros(437.87),
				candle.DollarsToMicros(428.43),
				candle.DollarsToMicros(428.35),
				candle.DollarsToMicros(432.5),
				candle.DollarsToMicros(443.66),
				candle.DollarsToMicros(455.72),
				candle.DollarsToMicros(454.49),
				candle.DollarsToMicros(452.08),
				candle.DollarsToMicros(452.73),
				candle.DollarsToMicros(461.91),
				candle.DollarsToMicros(463.58),
				candle.DollarsToMicros(461.14),
				candle.DollarsToMicros(452.08),
				candle.DollarsToMicros(442.66),
				candle.DollarsToMicros(428.91),
				candle.DollarsToMicros(429.79),
				candle.DollarsToMicros(431.99),
				candle.DollarsToMicros(427.72),
				candle.DollarsToMicros(423.2),
				candle.DollarsToMicros(426.21),
				candle.DollarsToMicros(426.98),
				candle.DollarsToMicros(435.69),
				candle.DollarsToMicros(434.33),
				candle.DollarsToMicros(429.8),
				candle.DollarsToMicros(419.85),
				candle.DollarsToMicros(426.24),
				candle.DollarsToMicros(402.8),
				candle.DollarsToMicros(392.05),
				candle.DollarsToMicros(390.53),
				candle.DollarsToMicros(398.67),
				candle.DollarsToMicros(406.13),
				candle.DollarsToMicros(405.46),
				candle.DollarsToMicros(408.38),
				candle.DollarsToMicros(417.2),
				candle.DollarsToMicros(430.12),
				candle.DollarsToMicros(442.78),
			},
			append(
				make([]ValidMicro, 33, 50),
				genValidMicro(-5108084),
				genValidMicro(-4527496),
				genValidMicro(-3387777),
				genValidMicro(-2592274),
				genValidMicro(-2250615),
				genValidMicro(-2552088),
				genValidMicro(-2192264),
				genValidMicro(-3335498),
				genValidMicro(-4543441),
				genValidMicro(-5129228),
				genValidMicro(-4666182),
				genValidMicro(-3602783),
				genValidMicro(-2729465),
				genValidMicro(-1785740),
				genValidMicro(-466764),
				genValidMicro(1280988),
				genValidMicro(3186354),
			),
		),
	)
}

func testStandardMACDFunc(closingPrices []int64, expected []ValidMicro) func(*testing.T) {
	return func(t *testing.T) {
		actual := StandardMACD(closingPrices)
		if !eqValidMicroSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

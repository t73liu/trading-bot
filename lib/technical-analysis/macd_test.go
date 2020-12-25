package analyze

import (
	"testing"
	"tradingbot/lib/utils"
)

func TestStandardMACD(t *testing.T) {
	t.Run(
		"Standard MACD not enough elements for any calculation",
		testStandardMACDFunc(
			[]int64{10, 10, 10, 10},
			make([]utils.MicroDollar, 4),
		),
	)
	t.Run(
		"Standard MACD enough elements for one calculation",
		testStandardMACDFunc(
			[]int64{
				utils.DollarsToMicros(459.99),
				utils.DollarsToMicros(448.85),
				utils.DollarsToMicros(446.06),
				utils.DollarsToMicros(450.81),
				utils.DollarsToMicros(442.8),
				utils.DollarsToMicros(448.97),
				utils.DollarsToMicros(444.57),
				utils.DollarsToMicros(441.4),
				utils.DollarsToMicros(430.47),
				utils.DollarsToMicros(420.05),
				utils.DollarsToMicros(431.14),
				utils.DollarsToMicros(425.66),
				utils.DollarsToMicros(430.58),
				utils.DollarsToMicros(431.72),
				utils.DollarsToMicros(437.87),
				utils.DollarsToMicros(428.43),
				utils.DollarsToMicros(428.35),
				utils.DollarsToMicros(432.5),
				utils.DollarsToMicros(443.66),
				utils.DollarsToMicros(455.72),
				utils.DollarsToMicros(454.49),
				utils.DollarsToMicros(452.08),
				utils.DollarsToMicros(452.73),
				utils.DollarsToMicros(461.91),
				utils.DollarsToMicros(463.58),
				utils.DollarsToMicros(461.14),
				utils.DollarsToMicros(452.08),
				utils.DollarsToMicros(442.66),
				utils.DollarsToMicros(428.91),
				utils.DollarsToMicros(429.79),
				utils.DollarsToMicros(431.99),
				utils.DollarsToMicros(427.72),
				utils.DollarsToMicros(423.2),
				utils.DollarsToMicros(426.21),
			},
			// Same as SMA for initial calc: (10 + 10 + 10 + 10 + 15) / 5 = 11
			append(
				make([]utils.MicroDollar, 33, 34),
				utils.NewMicroDollar(-5108084),
			),
		),
	)
	t.Run(
		"Standard MACD enough elements for multiple calculations",
		testStandardMACDFunc(
			[]int64{
				utils.DollarsToMicros(459.99),
				utils.DollarsToMicros(448.85),
				utils.DollarsToMicros(446.06),
				utils.DollarsToMicros(450.81),
				utils.DollarsToMicros(442.8),
				utils.DollarsToMicros(448.97),
				utils.DollarsToMicros(444.57),
				utils.DollarsToMicros(441.4),
				utils.DollarsToMicros(430.47),
				utils.DollarsToMicros(420.05),
				utils.DollarsToMicros(431.14),
				utils.DollarsToMicros(425.66),
				utils.DollarsToMicros(430.58),
				utils.DollarsToMicros(431.72),
				utils.DollarsToMicros(437.87),
				utils.DollarsToMicros(428.43),
				utils.DollarsToMicros(428.35),
				utils.DollarsToMicros(432.5),
				utils.DollarsToMicros(443.66),
				utils.DollarsToMicros(455.72),
				utils.DollarsToMicros(454.49),
				utils.DollarsToMicros(452.08),
				utils.DollarsToMicros(452.73),
				utils.DollarsToMicros(461.91),
				utils.DollarsToMicros(463.58),
				utils.DollarsToMicros(461.14),
				utils.DollarsToMicros(452.08),
				utils.DollarsToMicros(442.66),
				utils.DollarsToMicros(428.91),
				utils.DollarsToMicros(429.79),
				utils.DollarsToMicros(431.99),
				utils.DollarsToMicros(427.72),
				utils.DollarsToMicros(423.2),
				utils.DollarsToMicros(426.21),
				utils.DollarsToMicros(426.98),
				utils.DollarsToMicros(435.69),
				utils.DollarsToMicros(434.33),
				utils.DollarsToMicros(429.8),
				utils.DollarsToMicros(419.85),
				utils.DollarsToMicros(426.24),
				utils.DollarsToMicros(402.8),
				utils.DollarsToMicros(392.05),
				utils.DollarsToMicros(390.53),
				utils.DollarsToMicros(398.67),
				utils.DollarsToMicros(406.13),
				utils.DollarsToMicros(405.46),
				utils.DollarsToMicros(408.38),
				utils.DollarsToMicros(417.2),
				utils.DollarsToMicros(430.12),
				utils.DollarsToMicros(442.78),
			},
			append(
				make([]utils.MicroDollar, 33, 50),
				utils.NewMicroDollar(-5108084),
				utils.NewMicroDollar(-4527496),
				utils.NewMicroDollar(-3387777),
				utils.NewMicroDollar(-2592274),
				utils.NewMicroDollar(-2250615),
				utils.NewMicroDollar(-2552088),
				utils.NewMicroDollar(-2192264),
				utils.NewMicroDollar(-3335498),
				utils.NewMicroDollar(-4543441),
				utils.NewMicroDollar(-5129228),
				utils.NewMicroDollar(-4666182),
				utils.NewMicroDollar(-3602783),
				utils.NewMicroDollar(-2729465),
				utils.NewMicroDollar(-1785740),
				utils.NewMicroDollar(-466764),
				utils.NewMicroDollar(1280988),
				utils.NewMicroDollar(3186354),
			),
		),
	)
}

func testStandardMACDFunc(closingPrices []int64, expected []utils.MicroDollar) func(*testing.T) {
	return func(t *testing.T) {
		actual := StandardMACD(closingPrices)
		if !utils.EqMicroDollarSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

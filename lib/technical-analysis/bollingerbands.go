package analyze

import (
	"math"
	"tradingbot/lib/utils"
)

func BollingerBands(prices []int64, smaInterval int) []MicroDollarRange {
	if len(prices) < smaInterval {
		return make([]MicroDollarRange, len(prices))
	}
	smaValues := SMA(prices, smaInterval)
	standardDeviations := make([]utils.MicroDollar, 0, len(prices))
	for i, sma := range smaValues {
		if sma.Valid {
			var sum float64
			for j := 0; j < smaInterval; j++ {
				diff := utils.MicrosToDollars(prices[i-j] - sma.Value())
				sum += diff * diff
			}
			standardDeviation := math.Sqrt(sum / float64(smaInterval))
			standardDeviations = append(standardDeviations, utils.NewMicroDollar(utils.DollarsToMicros(standardDeviation)))
		} else {
			standardDeviations = append(standardDeviations, utils.MicroDollar{})
		}
	}
	bollingerBands := make([]MicroDollarRange, 0, len(prices))
	for i, sma := range smaValues {
		if sma.Valid {
			standardDeviation := standardDeviations[i].Value()
			bollingerBands = append(bollingerBands, MicroDollarRange{
				Valid: true,
				High:  sma.Value() + 2*standardDeviation,
				Mid:   sma.Value(),
				Low:   sma.Value() - 2*standardDeviation,
			})
		} else {
			bollingerBands = append(bollingerBands, MicroDollarRange{})
		}
	}
	return bollingerBands
}

package analyze

import (
	"math"
	"tradingbot/lib/candle"
)

func BollingerBands(prices []int64, smaInterval int) []ValidMicroRange {
	if len(prices) < smaInterval {
		return make([]ValidMicroRange, len(prices))
	}
	smaValues := SMA(prices, smaInterval)
	standardDeviations := make([]ValidMicro, 0, len(prices))
	for i, sma := range smaValues {
		if sma.Valid {
			var sum float64
			for j := 0; j < smaInterval; j++ {
				diff := candle.MicrosToDollars(prices[i-j] - sma.Value)
				sum += diff * diff
			}
			standardDeviation := math.Sqrt(sum / float64(smaInterval))
			standardDeviations = append(standardDeviations, genValidMicro(candle.DollarsToMicros(standardDeviation)))
		} else {
			standardDeviations = append(standardDeviations, ValidMicro{})
		}
	}
	bollingerBands := make([]ValidMicroRange, 0, len(prices))
	for i, sma := range smaValues {
		if sma.Valid {
			standardDeviation := standardDeviations[i].Value
			bollingerBands = append(bollingerBands, ValidMicroRange{
				Valid: true,
				High:  sma.Value + 2*standardDeviation,
				Mid:   sma.Value,
				Low:   sma.Value - 2*standardDeviation,
			})
		} else {
			bollingerBands = append(bollingerBands, ValidMicroRange{})
		}
	}
	return bollingerBands
}

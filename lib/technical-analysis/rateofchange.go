package analyze

import (
	"tradingbot/lib/candle"
	"tradingbot/lib/utils"
)

func SmoothedRateOfChange(values []int64, averageInterval, rateOfChangeInterval int) (results []ValidFloat) {
	ema := EMA(values, averageInterval)
	if rateOfChangeInterval+averageInterval > len(values) {
		results = make([]ValidFloat, len(values))
	} else {
		prevIndex := 0
		currIndex := rateOfChangeInterval
		results = make([]ValidFloat, rateOfChangeInterval+averageInterval-1, len(values))
		for currIndex < len(values) {
			prevVal := ema[prevIndex]
			currVal := ema[currIndex]
			if currVal.Valid && prevVal.Valid {
				diff := currVal.Value - prevVal.Value
				percentChange := candle.MicrosToDollars(diff) / candle.MicrosToDollars(prevVal.Value) * 100
				results = append(results, genValidFloat(utils.RoundToTwoDecimals(percentChange)))
			}
			prevIndex++
			currIndex++
		}
	}
	return results
}

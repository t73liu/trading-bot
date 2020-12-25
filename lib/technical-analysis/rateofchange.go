package analyze

import (
	"tradingbot/lib/utils"
)

func SmoothedRateOfChange(values []int64, averageInterval, rateOfChangeInterval int) (results []utils.NullFloat64) {
	ema := EMA(values, averageInterval)
	if rateOfChangeInterval+averageInterval > len(values) {
		results = make([]utils.NullFloat64, len(values))
	} else {
		prevIndex := 0
		currIndex := rateOfChangeInterval
		results = make([]utils.NullFloat64, rateOfChangeInterval+averageInterval-1, len(values))
		for currIndex < len(values) {
			prevVal := ema[prevIndex]
			currVal := ema[currIndex]
			if currVal.Valid && prevVal.Valid {
				diff := currVal.Value() - prevVal.Value()
				percentChange := utils.MicrosToDollars(diff) / utils.MicrosToDollars(prevVal.Value()) * 100
				results = append(results, utils.NewNullFloat64(utils.RoundToTwoDecimals(percentChange)))
			}
			prevIndex++
			currIndex++
		}
	}
	return results
}

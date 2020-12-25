package analyze

import "tradingbot/lib/utils"

// Simple Moving Average
func SMA(values []int64, interval int) (results []utils.MicroDollar) {
	if len(values) >= interval && interval > 2 {
		results = make([]utils.MicroDollar, interval-1, len(values))
		formattedInterval := int64(interval)
		sum := calcSum(values, 0, interval-1)
		sma := sum / formattedInterval
		results = append(results, utils.NewMicroDollar(sma))
		laggedIndex := 0
		currentIndex := interval
		for currentIndex < len(values) {
			sum = sum + values[currentIndex] - values[laggedIndex]
			sma = sum / formattedInterval
			results = append(results, utils.NewMicroDollar(sma))
			laggedIndex++
			currentIndex++
		}
	} else {
		results = make([]utils.MicroDollar, len(values))
	}
	return results
}

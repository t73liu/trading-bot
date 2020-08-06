package analyze

import "tradingbot/lib/candle"

func getGains(values []int64) []int64 {
	gains := make([]int64, len(values))
	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff > 0 {
			gains[i] = diff
		}
	}
	return gains
}

func getAbsLosses(values []int64) []int64 {
	losses := make([]int64, len(values))
	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff < 0 {
			losses[i] = -1 * diff
		}
	}
	return losses
}

// RSI using Wilder's Smoothing Method
func RSI(values []int64, interval int) (results []ValidFloat) {
	if len(values) >= interval+1 && interval > 2 {
		results = make([]ValidFloat, interval, len(values))
		formattedInterval := int64(interval)
		gains := getGains(values)
		losses := getAbsLosses(values)
		averageGain := calcAverage(gains, 1, interval)
		averageLoss := calcAverage(losses, 1, interval)
		results = append(results, genValidFloat(calcRSI(averageGain, averageLoss)))
		for i := interval + 1; i < len(values); i++ {
			averageGain = ((formattedInterval-1)*averageGain + gains[i]) / formattedInterval
			averageLoss = ((formattedInterval-1)*averageLoss + losses[i]) / formattedInterval
			results = append(results, genValidFloat(calcRSI(averageGain, averageLoss)))
		}
	} else {
		size := interval
		if len(values) < interval {
			size = len(values)
		}
		results = make([]ValidFloat, size)
	}
	return results
}

func calcRSI(averageGain, averageLoss int64) float64 {
	if averageGain == 0 {
		return 0
	}
	if averageLoss == 0 {
		return 100
	}
	relativeStrength := candle.MicrosToDollars(averageGain) / candle.MicrosToDollars(averageLoss)
	rsi := 100 - (100 / (1 + relativeStrength))
	return RoundToTwoDecimals(rsi)
}

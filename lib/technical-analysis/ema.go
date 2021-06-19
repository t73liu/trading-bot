package analyze

import "github.com/t73liu/tradingbot/lib/utils"

// Exponential Moving Average
func EMA(values []int64, interval int) (results []utils.MicroDollar) {
	if len(values) >= interval && interval > 2 {
		results = make([]utils.MicroDollar, interval-1, len(values))
		ema := calcAverage(values, 0, interval-1)
		results = append(results, utils.NewMicroDollar(ema))
		multiplier := 2 / float64(interval+1)
		for i := interval; i < len(values); i++ {
			ema = int64(float64(values[i]-ema)*multiplier) + ema
			results = append(results, utils.NewMicroDollar(ema))
		}
	} else {
		results = make([]utils.MicroDollar, len(values))
	}
	return results
}

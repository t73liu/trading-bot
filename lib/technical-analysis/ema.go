package analyze

// Exponential Moving Average
func EMA(values []int64, interval int) (results []ValidMicro) {
	if len(values) >= interval && interval > 2 {
		results = make([]ValidMicro, interval-1, len(values))
		ema := calcAverage(values, 0, interval-1)
		results = append(results, genValidMicro(ema))
		multiplier := 2 / float64(interval+1)
		for i := interval; i < len(values); i++ {
			ema = int64(float64(values[i]-ema)*multiplier) + ema
			results = append(results, genValidMicro(ema))
		}
	} else {
		results = make([]ValidMicro, len(values))
	}
	return results
}

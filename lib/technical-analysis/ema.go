package analyze

// Exponential Moving Average
func EMA(candles []Candle, interval int) (results []ValidCalc) {
	if len(candles) >= interval && interval > 1 {
		results = make([]ValidCalc, interval-1)
		formattedInterval := int64(interval)
		var sum int64
		for i := 0; i < interval; i++ {
			sum += candles[i].Close
		}
		value := sum / formattedInterval
		results = append(results, genValidCalc(value))
		multiplier := 2 / float64(interval+1)
		for i := interval; i < len(candles); i++ {
			value = int64(float64(candles[i].Close-value)*multiplier) + value
			results = append(results, genValidCalc(value))
		}
	} else {
		results = make([]ValidCalc, len(candles))
	}
	return results
}

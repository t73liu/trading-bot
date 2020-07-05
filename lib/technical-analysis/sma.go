package analyze

// Simple Moving Average
func SMA(candles []Candle, interval int) (results []ValidCalc) {
	if len(candles) >= interval && interval > 1 {
		results = make([]ValidCalc, interval-1)
		formattedInterval := int64(interval)
		var sum int64
		for i := 0; i < interval; i++ {
			sum += candles[i].Close
		}
		value := sum / formattedInterval
		results = append(results, genValidCalc(value))
		laggedIndex := 0
		currentIndex := interval
		for currentIndex < len(candles) {
			sum = sum + candles[currentIndex].Close - candles[laggedIndex].Close
			value = sum / formattedInterval
			results = append(results, genValidCalc(value))
			laggedIndex++
			currentIndex++
		}
	} else {
		results = make([]ValidCalc, len(candles))
	}
	return results
}

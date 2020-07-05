package analyze

// Simple Moving Average
func SMA(values []int64, interval int) (results []ValidMicro) {
	if len(values) >= interval && interval > 2 {
		results = make([]ValidMicro, interval-1, len(values))
		formattedInterval := int64(interval)
		sum := calcSum(values, 0, interval-1)
		sma := sum / formattedInterval
		results = append(results, genValidMicro(sma))
		laggedIndex := 0
		currentIndex := interval
		for currentIndex < len(values) {
			sum = sum + values[currentIndex] - values[laggedIndex]
			sma = sum / formattedInterval
			results = append(results, genValidMicro(sma))
			laggedIndex++
			currentIndex++
		}
	} else {
		results = make([]ValidMicro, len(values))
	}
	return results
}

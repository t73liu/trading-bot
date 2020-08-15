package analyze

func StandardMACD(values []int64) []ValidMicro {
	return MACD(values, 12, 26, 9)
}

func MACD(values []int64, fastEmaInterval, slowEmaInterval, signalInterval int) (results []ValidMicro) {
	requiredInterval := slowEmaInterval + signalInterval - 1
	if len(values) >= requiredInterval && requiredInterval > 5 {
		results = make([]ValidMicro, requiredInterval-1, len(values))
		fastEMAs := EMA(values, fastEmaInterval)
		slowEMAs := EMA(values, slowEmaInterval)
		macdLine := make([]int64, 0, len(values)-slowEmaInterval)
		for i := slowEmaInterval - 1; i < len(values); i++ {
			macdLine = append(macdLine, fastEMAs[i].Value-slowEMAs[i].Value)
		}
		signalEMAs := EMA(macdLine, 9)
		for i, signal := range signalEMAs {
			if signal.Valid {
				results = append(results, genValidMicro(macdLine[i]-signal.Value))
			}
		}
	} else {
		results = make([]ValidMicro, len(values))
	}
	return results
}

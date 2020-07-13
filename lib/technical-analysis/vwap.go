package analyze

// Volume Weighted Average Price
func VWAP(candles []Candle) (results []ValidMicro) {
	if len(candles) > 0 {
		results = make([]ValidMicro, 0, len(candles))
		var totalTypicalPrice int64
		var totalVolume int64
		for _, candle := range candles {
			totalTypicalPrice += (candle.High + candle.Low + candle.Close) / 3 * candle.Volume
			totalVolume += candle.Volume
			results = append(results, genValidMicro(totalTypicalPrice/totalVolume))
		}
	}
	return results
}

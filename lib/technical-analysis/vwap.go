package analyze

import "tradingbot/lib/candle"

// Volume Weighted Average Price
func VWAP(candles []candle.Candle) (results []ValidMicro) {
	if len(candles) > 0 {
		results = make([]ValidMicro, 0, len(candles))
		var totalTypicalPrice int64
		var totalVolume int64
		for _, c := range candles {
			totalTypicalPrice += (c.HighMicros + c.LowMicros + c.CloseMicros) / 3 * c.Volume
			totalVolume += c.Volume
			results = append(results, genValidMicro(totalTypicalPrice/totalVolume))
		}
	}
	return results
}

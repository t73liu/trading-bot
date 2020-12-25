package analyze

import (
	"tradingbot/lib/candle"
	"tradingbot/lib/utils"
)

// Volume Weighted Average Price
func VWAP(candles []candle.Candle) (results []utils.MicroDollar) {
	if len(candles) > 0 {
		results = make([]utils.MicroDollar, 0, len(candles))
		var totalTypicalPrice int64
		var totalVolume int64
		for _, c := range candles {
			totalTypicalPrice += (c.HighMicros + c.LowMicros + c.CloseMicros) / 3 * c.Volume
			totalVolume += c.Volume
			results = append(results, utils.NewMicroDollar(totalTypicalPrice/totalVolume))
		}
	}
	return results
}

package analyze

import (
	"tradingbot/lib/candle"
	"tradingbot/lib/utils"
)

// TTM Squeeze returns true when Bollinger Bands are between Keltner Channels
func TTMSqueeze(candles []candle.Candle) []utils.NullBool {
	if len(candles) < 20 {
		return make([]utils.NullBool, len(candles))
	}
	output := make([]utils.NullBool, 0, len(candles))
	keltnerChannelValues := KeltnerChannels(candles, 20, 10)
	bollingerBandValues := BollingerBands(candle.GetClosingPrices(candles), 20)
	for i, kc := range keltnerChannelValues {
		bb := bollingerBandValues[i]
		if bb.Valid && kc.Valid {
			output = append(output, utils.NewNullBool(bb.Low > kc.Low && bb.High < kc.High))
		} else {
			output = append(output, utils.NullBool{})
		}
	}
	return output
}

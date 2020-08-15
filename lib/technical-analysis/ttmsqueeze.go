package analyze

import "tradingbot/lib/candle"

// TTM Squeeze returns true when Bollinger Bands are between Keltner Channels
func TTMSqueeze(candles []candle.Candle) []ValidBool {
	if len(candles) < 20 {
		return make([]ValidBool, len(candles))
	}
	output := make([]ValidBool, 0, len(candles))
	keltnerChannelValues := KeltnerChannels(candles, 20, 10)
	bollingerBandValues := BollingerBands(candle.GetClosingPrices(candles), 20)
	for i, kc := range keltnerChannelValues {
		bb := bollingerBandValues[i]
		if bb.Valid && kc.Valid {
			output = append(output, ValidBool{
				Valid: true,
				Value: bb.Low > kc.Low && bb.High < kc.High,
			})
		} else {
			output = append(output, ValidBool{})
		}
	}
	return output
}

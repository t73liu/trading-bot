package analyze

import "github.com/t73liu/tradingbot/lib/candle"

// Keltner Channels
func KeltnerChannels(candles []candle.Candle, emaInterval int, atrInterval int) []MicroDollarRange {
	if len(candles) < emaInterval || len(candles) < atrInterval {
		return make([]MicroDollarRange, len(candles))
	}
	atrValues := ATR(candles, atrInterval)
	emaValues := EMA(candle.GetClosingPrices(candles), emaInterval)
	keltnerValues := make([]MicroDollarRange, 0, len(candles))
	for i, ema := range emaValues {
		atr := atrValues[i]
		if atr.Valid && ema.Valid {
			keltnerValues = append(keltnerValues, MicroDollarRange{
				High:  ema.Value() + 2*atr.Value(),
				Mid:   ema.Value(),
				Low:   ema.Value() - 2*atr.Value(),
				Valid: true,
			})
		} else {
			keltnerValues = append(keltnerValues, MicroDollarRange{})
		}
	}
	return keltnerValues
}

package analyze

import (
	"tradingbot/lib/candle"
	"tradingbot/lib/utils"
)

// Average True Range
func ATR(candles []candle.Candle, interval int) []utils.MicroDollar {
	if len(candles) < interval {
		return make([]utils.MicroDollar, len(candles))
	}
	trueRanges := make([]int64, 0, len(candles))
	var prev candle.Candle
	for _, c := range candles {
		trueRanges = append(trueRanges, calcTrueRange(c, prev))
		prev = c
	}
	output := make([]utils.MicroDollar, interval-1, len(candles))
	var sum, atr int64
	for i, trueRange := range trueRanges {
		if i < interval {
			sum += trueRange
			if i == interval-1 {
				atr = sum / int64(interval)
				output = append(output, utils.NewMicroDollar(atr))
			}
		} else {
			atr = ((atr * int64(interval-1)) + trueRange) / int64(interval)
			output = append(output, utils.NewMicroDollar(atr))
		}
	}
	return output
}

func calcTrueRange(curr candle.Candle, prev candle.Candle) int64 {
	highLow := curr.HighMicros - curr.LowMicros
	if prev.IsZero() {
		return highLow
	}
	highPrevClose := utils.Abs(curr.HighMicros - prev.CloseMicros)
	lowPrevClose := utils.Abs(curr.LowMicros - prev.CloseMicros)
	return utils.Max(highLow, highPrevClose, lowPrevClose)
}

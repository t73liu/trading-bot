package analyze

import (
	"time"
)

type Candle struct {
	OpenedAt time.Time
	Volume   int64
	Open     int64
	High     int64
	Low      int64
	Close    int64
}

// Compress minute candles to hourly, daily, etc.
func CompressCandles(candles []Candle, timeInterval int, timeUnit string) []Candle {
	if len(candles) == 0 || timeInterval == 1 && timeUnit == "minute" {
		return candles
	}
	newCandles := make([]Candle, 0, len(candles)/timeInterval)
	count := 0
	// TODO need to respect day separation and non-zero starts
	for _, candle := range candles {
		if count == timeInterval-1 {
			newCandles = append(newCandles, candle)
			count = 0
		} else {
			count++
		}
	}
	return newCandles
}

// Fill in gaps for consecutive periods
func FillMinuteCandles(candles []Candle) []Candle {
	if len(candles) == 0 {
		return candles
	}
	filledCandles := make([]Candle, 0, len(candles))
	var prevCandle Candle
	var prevTime time.Time
	for i, candle := range candles {
		currentTime := candle.OpenedAt
		if i > 0 && prevTime.Day() == currentTime.Day() {
			minutesDiff := int(currentTime.Sub(prevTime).Minutes())
			for i := minutesDiff; i > 1; i-- {
				backfilledTime := currentTime.Add(-1 * time.Minute * time.Duration(i-1))
				filledCandles = append(filledCandles, GenPlaceholderCandle(prevCandle, backfilledTime))
			}
		}
		prevCandle = candle
		prevTime = prevCandle.OpenedAt
		filledCandles = append(filledCandles, candle)
	}
	return filledCandles
}

func GenPlaceholderCandle(candle Candle, openedAt time.Time) Candle {
	return Candle{
		OpenedAt: openedAt,
		Volume:   0,
		Open:     candle.Close,
		High:     candle.Close,
		Low:      candle.Close,
		Close:    candle.Close,
	}
}

func eqCandleSlice(expected, actual []Candle) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i := range expected {
		if expected[i] != actual[i] {
			return false
		}
	}
	return true
}

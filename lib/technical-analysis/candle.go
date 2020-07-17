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
func CompressCandles(candles []Candle, timeInterval int, timeUnit string, loc *time.Location) []Candle {
	if len(candles) == 0 || timeInterval == 1 && timeUnit == "minute" {
		return candles
	}
	newCandles := make([]Candle, 0, len(candles)/timeInterval)
	var prevTimeBucket time.Time
	for _, candle := range candles {
		currentTimeBucket := getKey(candle, timeInterval, timeUnit, loc)
		if prevTimeBucket.IsZero() || currentTimeBucket != prevTimeBucket {
			prevTimeBucket = currentTimeBucket
			candle.OpenedAt = currentTimeBucket
			newCandles = append(newCandles, candle)
		} else {
			lastIndex := len(newCandles) - 1
			newCandles[lastIndex].Volume += candle.Volume
			if candle.High > newCandles[lastIndex].High {
				newCandles[lastIndex].High = candle.High
			}
			if candle.Low < newCandles[lastIndex].Low {
				newCandles[lastIndex].Low = candle.Low
			}
			newCandles[lastIndex].Close = candle.Close
		}
	}
	return newCandles
}

func getKey(candle Candle, timeInterval int, timeUnit string, loc *time.Location) time.Time {
	openedAt := candle.OpenedAt
	switch timeUnit {
	case "minute":
		return GetMinuteBucket(openedAt, loc, timeInterval)
	case "hour":
		return GetHourBucket(openedAt, loc, timeInterval)
	case "day":
		return GetMidnight(openedAt, loc)
	case "week":
		return GetStartOfWeek(openedAt, loc)
	case "month":
		return GetStartOfMonth(openedAt, loc)
	default:
		return openedAt
	}
}

// TODO Move to utils package
func GetMinuteBucket(moment time.Time, loc *time.Location, interval int) time.Time {
	year, month, day := moment.Date()
	hour, minute, _ := moment.Clock()
	return time.Date(year, month, day, hour, minute/interval, 0, 0, loc)
}

func GetHourBucket(moment time.Time, loc *time.Location, interval int) time.Time {
	year, month, day := moment.Date()
	hour, _, _ := moment.Clock()
	return time.Date(year, month, day, hour/interval, 0, 0, 0, loc)
}

func GetMidnight(moment time.Time, loc *time.Location) time.Time {
	year, month, day := moment.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}

func GetStartOfWeek(moment time.Time, loc *time.Location) time.Time {
	current := moment
	for current.Weekday() != time.Monday {
		current = current.AddDate(0, 0, -1)
	}
	return GetMidnight(current, loc)
}

func GetStartOfMonth(moment time.Time, loc *time.Location) time.Time {
	year, month, _ := moment.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, loc)
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

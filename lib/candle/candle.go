package candle

import (
	"encoding/json"
	"errors"
	"time"
	"tradingbot/lib/utils"
)

type JSONCandle struct {
	OpenedAt time.Time `json:"openedAt"`
	Volume   int64     `json:"volume"`
	Open     float64   `json:"open"`
	High     float64   `json:"high"`
	Low      float64   `json:"low"`
	Close    float64   `json:"close"`
}

type Candle struct {
	OpenedAt    time.Time
	Volume      int64
	OpenMicros  int64
	HighMicros  int64
	LowMicros   int64
	CloseMicros int64
}

func (c Candle) MarshalJSON() ([]byte, error) {
	return json.Marshal(JSONCandle{
		OpenedAt: c.OpenedAt,
		Volume:   c.Volume,
		Open:     MicrosToDollars(c.OpenMicros),
		High:     MicrosToDollars(c.HighMicros),
		Low:      MicrosToDollars(c.LowMicros),
		Close:    MicrosToDollars(c.CloseMicros),
	})
}

const million = 1000000

func DollarsToMicros(dollars float64) int64 {
	return int64(dollars * million)
}

func MicrosToDollars(micros int64) float64 {
	return float64(micros) / million
}

// Assuming location = "America/New_York"
func FilterTradingHourCandles(candles []Candle) []Candle {
	filtered := make([]Candle, 0, len(candles))
	for _, candle := range candles {
		if utils.IsWithinNYSETradingHours(candle.OpenedAt) {
			filtered = append(filtered, candle)
		}
	}
	return filtered
}

// Compress minute candles to hourly, daily, etc.
func CompressCandles(candles []Candle, timeInterval uint, timeUnit string, loc *time.Location) ([]Candle, error) {
	if len(candles) == 0 || timeInterval == 1 && timeUnit == "minute" {
		return candles, nil
	}
	if timeInterval != 1 && timeUnit != "minute" && timeUnit != "hour" {
		return nil, errors.New("only minute/hour can specify time interval greater than 1")
	}
	if timeInterval > 30 && timeUnit == "minute" {
		return nil, errors.New("minute should have a max time interval of 30")
	}
	if timeInterval > 12 && timeUnit == "hour" {
		return nil, errors.New("hour should have a max time interval of 12")
	}
	newCandles := make([]Candle, 0, len(candles)/int(timeInterval))
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
			if candle.HighMicros > newCandles[lastIndex].HighMicros {
				newCandles[lastIndex].HighMicros = candle.HighMicros
			}
			if candle.LowMicros < newCandles[lastIndex].LowMicros {
				newCandles[lastIndex].LowMicros = candle.LowMicros
			}
			newCandles[lastIndex].CloseMicros = candle.CloseMicros
		}
	}
	return newCandles, nil
}

func getKey(candle Candle, timeInterval uint, timeUnit string, loc *time.Location) time.Time {
	openedAt := candle.OpenedAt
	switch timeUnit {
	case "minute":
		return utils.GetMinuteBucket(openedAt, loc, timeInterval)
	case "hour":
		return utils.GetHourBucket(openedAt, loc, timeInterval)
	case "day":
		return utils.GetMidnight(openedAt, loc)
	case "week":
		return utils.GetStartOfWeek(openedAt, loc)
	case "month":
		return utils.GetStartOfMonth(openedAt, loc)
	default:
		return openedAt
	}
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
		OpenedAt:    openedAt,
		Volume:      0,
		OpenMicros:  candle.CloseMicros,
		HighMicros:  candle.CloseMicros,
		LowMicros:   candle.CloseMicros,
		CloseMicros: candle.CloseMicros,
	}
}

func GetClosingPrices(candles []Candle) (closes []int64) {
	for _, candle := range candles {
		closes = append(closes, candle.CloseMicros)
	}
	return closes
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

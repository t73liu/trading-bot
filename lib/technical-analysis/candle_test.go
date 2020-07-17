package analyze

import (
	"testing"
	"time"
)

// TODO complete test
func TestCompressCandles(t *testing.T) {
	now := time.Date(2020, 7, 1, 9, 1, 0, 0, time.UTC)
	//tomorrow := addMinutes(now.AddDate(0, 0, 1), 13)
	//nextWeek := addMinutes(now.AddDate(0, 0, 7), 46)
	t.Run(
		"No candles provided (1-minute)",
		testCompressCandlesFunc([]Candle{}, 1, "minute", []Candle{}),
	)
	t.Run(
		"nil provided (1-hour)",
		testCompressCandlesFunc(nil, 1, "hour", nil),
	)
	t.Run(
		"No candles provided (30-minute)",
		testCompressCandlesFunc([]Candle{}, 30, "minute", []Candle{}),
	)
	t.Run(
		"nil provided (1-day)",
		testCompressCandlesFunc(nil, 1, "day", nil),
	)
	t.Run(
		"nil provided (1-day)",
		testCompressCandlesFunc(
			[]Candle{
				{
					OpenedAt: now,
					Volume:   50,
					Open:     11,
					High:     33,
					Low:      7,
					Close:    34,
				},
				{
					OpenedAt: addMinutes(now, 60),
					Volume:   510,
					Open:     344,
					High:     588,
					Low:      77,
					Close:    2,
				},
			},
			1,
			"day",
			[]Candle{
				{
					OpenedAt: time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
					Open:     11,
					High:     588,
					Close:    2,
					Low:      7,
					Volume:   560,
				},
			},
		),
	)
}

func TestGetMidnight(t *testing.T) {
	midnight := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
	midnightResult := GetMidnight(midnight, time.UTC)
	if midnightResult != midnight {
		t.Errorf("Midnight input failed, expected %s, got %s\n", midnight, midnightResult)
	}
	quarterPastNine := time.Date(2020, 7, 1, 9, 15, 0, 0, time.UTC)
	quarterPastNineResult := GetMidnight(quarterPastNine, time.UTC)
	if quarterPastNineResult != midnight {
		t.Errorf("9:15 input failed, expected %s, got %s\n", midnight, quarterPastNineResult)
	}
}

func TestGenPlaceholderCandle(t *testing.T) {
	now := time.Now()
	emptyResult := GenPlaceholderCandle(Candle{}, now)
	expectedEmptyResult := Candle{OpenedAt: now}
	if emptyResult != expectedEmptyResult {
		t.Errorf("Empty result failed, expected %+v, got %+v\n", expectedEmptyResult, emptyResult)
	}

	result := GenPlaceholderCandle(
		Candle{
			OpenedAt: time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
			Volume:   135,
			Open:     456,
			High:     789,
			Low:      753,
			Close:    159,
		},
		now,
	)
	expectedResult := Candle{
		OpenedAt: now,
		Volume:   0,
		Open:     159,
		High:     159,
		Low:      159,
		Close:    159,
	}
	if result != expectedResult {
		t.Errorf("Result failed, expected %+v, got %+v\n", expectedEmptyResult, emptyResult)
	}
}

func TestFillMinuteCandles(t *testing.T) {
	now := time.Date(2020, 7, 1, 9, 1, 0, 0, time.UTC)
	tomorrow := addMinutes(now.AddDate(0, 0, 1), 13)
	nextWeek := addMinutes(now.AddDate(0, 0, 7), 67)
	t.Run(
		"No candles provided",
		testFillMinuteCandlesFunc([]Candle{}, []Candle{}),
	)
	t.Run(
		"nil provided",
		testFillMinuteCandlesFunc(nil, nil),
	)
	t.Run(
		"Single day with no candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Low: 951, Volume: 852},
			},
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Low: 951, Volume: 852},
			},
		),
	)
	t.Run(
		"Single day with single candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 3), Low: 951, Volume: 852},
				{OpenedAt: addMinutes(now, 5), Low: 888, Volume: 777},
			},
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Open: 852, High: 852, Low: 852, Close: 852},
				{OpenedAt: addMinutes(now, 3), Low: 951, Volume: 852},
				{OpenedAt: addMinutes(now, 4)},
				{OpenedAt: addMinutes(now, 5), Low: 888, Volume: 777},
			},
		),
	)
	t.Run(
		"Single day with multiple candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 3), Low: 951, Volume: 852},
				{OpenedAt: addMinutes(now, 5), Low: 888, Volume: 777},
			},
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Open: 789, High: 789, Low: 789, Close: 789},
				{OpenedAt: addMinutes(now, 2), Open: 789, High: 789, Low: 789, Close: 789},
				{OpenedAt: addMinutes(now, 3), Low: 951, Volume: 852},
				{OpenedAt: addMinutes(now, 4)},
				{OpenedAt: addMinutes(now, 5), Low: 888, Volume: 777},
			},
		),
	)
	t.Run(
		"Multiple days with no candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Low: 951, Volume: 852},
				{OpenedAt: tomorrow, Close: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), Open: 788, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 2), Low: 987, Volume: 1},
				{OpenedAt: nextWeek, High: 999, Volume: 2},
			},
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Low: 951, Volume: 852},
				{OpenedAt: tomorrow, Close: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), Open: 788, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 2), Low: 987, Volume: 1},
				{OpenedAt: nextWeek, High: 999, Volume: 2},
			},
		),
	)
	t.Run(
		"Multiple days with single candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Low: 951, Volume: 852},
				{OpenedAt: tomorrow, Close: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 2), Low: 987, Volume: 1},
				{OpenedAt: nextWeek, High: 999, Volume: 2},
			},
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Low: 951, Volume: 852},
				{OpenedAt: tomorrow, Close: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), Open: 456, High: 456, Low: 456, Close: 456},
				{OpenedAt: addMinutes(tomorrow, 2), Low: 987, Volume: 1},
				{OpenedAt: nextWeek, High: 999, Volume: 2},
			},
		),
	)
	t.Run(
		"Multiple days with multiple candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Low: 951, Volume: 852},
				{OpenedAt: tomorrow, Close: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), Low: 777, Close: 774, Volume: 10},
				{OpenedAt: addMinutes(tomorrow, 5), Low: 987, Volume: 1},
				{OpenedAt: nextWeek, High: 999, Volume: 2},
				{OpenedAt: addMinutes(nextWeek, 2), High: 77, Volume: 50},
			},
			[]Candle{
				{OpenedAt: now, Open: 159, High: 123, Low: 456, Close: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), Close: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), Low: 951, Volume: 852},
				{OpenedAt: tomorrow, Close: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), Low: 777, Close: 774, Volume: 10},
				{OpenedAt: addMinutes(tomorrow, 2), Open: 774, High: 774, Low: 774, Close: 774},
				{OpenedAt: addMinutes(tomorrow, 3), Open: 774, High: 774, Low: 774, Close: 774},
				{OpenedAt: addMinutes(tomorrow, 4), Open: 774, High: 774, Low: 774, Close: 774},
				{OpenedAt: addMinutes(tomorrow, 5), Low: 987, Volume: 1},
				{OpenedAt: nextWeek, High: 999, Volume: 2},
				{OpenedAt: addMinutes(nextWeek, 1)},
				{OpenedAt: addMinutes(nextWeek, 2), High: 77, Volume: 50},
			},
		),
	)
}

func testFillMinuteCandlesFunc(candles []Candle, expected []Candle) func(*testing.T) {
	return func(t *testing.T) {
		actual := FillMinuteCandles(candles)
		if !eqCandleSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

func testCompressCandlesFunc(candles []Candle, timeInterval int, timeUnit string, expected []Candle) func(*testing.T) {
	return func(t *testing.T) {
		actual := CompressCandles(candles, timeInterval, timeUnit, time.UTC)
		if !eqCandleSlice(expected, actual) {
			t.Errorf("\nExpected: %v\nActual: %v", expected, actual)
		}
	}
}

func addMinutes(moment time.Time, minutes int) time.Time {
	return moment.Add(time.Duration(minutes) * time.Minute)
}

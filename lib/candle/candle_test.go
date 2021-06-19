package candle

import (
	"testing"
	"time"

	"github.com/t73liu/tradingbot/lib/utils"
)

func TestCompressCandles(t *testing.T) {
	now := time.Date(2020, 7, 1, 9, 1, 0, 0, time.UTC)
	tomorrow := addMinutes(now.AddDate(0, 0, 1), 13)
	nextWeek := addMinutes(now.AddDate(0, 0, 7), 46)
	testCandles := []Candle{
		{
			OpenedAt:    now,
			Volume:      50,
			OpenMicros:  11,
			HighMicros:  34,
			LowMicros:   7,
			CloseMicros: 33,
		},
		{
			OpenedAt:    addMinutes(now, 60),
			Volume:      510,
			OpenMicros:  344,
			HighMicros:  588,
			LowMicros:   77,
			CloseMicros: 79,
		},
		{
			OpenedAt:    addMinutes(now, 62),
			Volume:      223,
			OpenMicros:  321,
			HighMicros:  455,
			LowMicros:   78,
			CloseMicros: 79,
		},
		{
			OpenedAt:    addMinutes(now, 450),
			Volume:      112,
			OpenMicros:  34,
			HighMicros:  88,
			LowMicros:   70,
			CloseMicros: 72,
		},
		{
			OpenedAt:    tomorrow,
			Volume:      85,
			OpenMicros:  98,
			HighMicros:  123,
			LowMicros:   98,
			CloseMicros: 98,
		},
		{
			OpenedAt:    addMinutes(tomorrow, 360),
			Volume:      775,
			OpenMicros:  951,
			HighMicros:  951,
			LowMicros:   515,
			CloseMicros: 723,
		},
		{
			OpenedAt:    addMinutes(nextWeek, 25),
			Volume:      975,
			OpenMicros:  289,
			HighMicros:  321,
			LowMicros:   77,
			CloseMicros: 289,
		},
	}
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
		"multiple candles for single day (1-day)",
		testCompressCandlesFunc(
			[]Candle{
				{
					OpenedAt:    now,
					Volume:      50,
					OpenMicros:  11,
					HighMicros:  34,
					LowMicros:   7,
					CloseMicros: 33,
				},
				{
					OpenedAt:    addMinutes(now, 60),
					Volume:      510,
					OpenMicros:  344,
					HighMicros:  588,
					LowMicros:   77,
					CloseMicros: 79,
				},
				{
					OpenedAt:    addMinutes(now, 450),
					Volume:      112,
					OpenMicros:  34,
					HighMicros:  88,
					LowMicros:   70,
					CloseMicros: 72,
				},
			},
			1,
			"day",
			[]Candle{
				{
					OpenedAt: time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
					// 50 + 510 + 112 = 672
					Volume:      672,
					OpenMicros:  11,
					HighMicros:  588,
					LowMicros:   7,
					CloseMicros: 72,
				},
			},
		),
	)
	t.Run(
		"multiple candles for multiple days provided (1-day)",
		testCompressCandlesFunc(
			testCandles,
			1,
			"day",
			[]Candle{
				{
					OpenedAt: time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
					// 50 + 510 + 223 + 112 = 895
					Volume:      895,
					OpenMicros:  11,
					HighMicros:  588,
					LowMicros:   7,
					CloseMicros: 72,
				},
				{
					OpenedAt: time.Date(2020, 7, 2, 0, 0, 0, 0, time.UTC),
					// 85 + 775 = 860
					Volume:      860,
					OpenMicros:  98,
					HighMicros:  951,
					LowMicros:   98,
					CloseMicros: 723,
				},
				{
					OpenedAt:    time.Date(2020, 7, 8, 0, 0, 0, 0, time.UTC),
					Volume:      975,
					OpenMicros:  289,
					HighMicros:  321,
					LowMicros:   77,
					CloseMicros: 289,
				},
			},
		),
	)
	t.Run(
		"multiple candles for multiple days provided (1-week)",
		testCompressCandlesFunc(
			testCandles,
			1,
			"week",
			[]Candle{
				{
					OpenedAt: time.Date(2020, 6, 29, 0, 0, 0, 0, time.UTC),
					// 50 + 510 + 223 + 112 + 85 + 775 = 1755
					Volume:      1755,
					OpenMicros:  11,
					HighMicros:  951,
					LowMicros:   7,
					CloseMicros: 723,
				},
				{
					OpenedAt:    time.Date(2020, 7, 6, 0, 0, 0, 0, time.UTC),
					Volume:      975,
					OpenMicros:  289,
					HighMicros:  321,
					LowMicros:   77,
					CloseMicros: 289,
				},
			},
		),
	)
	t.Run(
		"multiple candles for multiple days provided (1-month)",
		testCompressCandlesFunc(
			testCandles,
			1,
			"month",
			[]Candle{
				{
					OpenedAt: time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
					// 50 + 510 + 223 + 112 + 85 + 775 + 975 = 2730
					Volume:      2730,
					OpenMicros:  11,
					HighMicros:  951,
					LowMicros:   7,
					CloseMicros: 289,
				},
			},
		),
	)
	t.Run(
		"multiple candles for multiple days provided (5-minute)",
		testCompressCandlesFunc(
			testCandles,
			5,
			"minute",
			[]Candle{
				{
					OpenedAt:    time.Date(2020, 7, 1, 9, 0, 0, 0, time.UTC),
					Volume:      50,
					OpenMicros:  11,
					HighMicros:  34,
					LowMicros:   7,
					CloseMicros: 33,
				},
				{
					OpenedAt: time.Date(2020, 7, 1, 10, 0, 0, 0, time.UTC),
					// 510 + 223 = 733
					Volume:      733,
					OpenMicros:  344,
					HighMicros:  588,
					LowMicros:   77,
					CloseMicros: 79,
				},
				{
					OpenedAt:    time.Date(2020, 7, 1, 16, 30, 0, 0, time.UTC),
					Volume:      112,
					OpenMicros:  34,
					HighMicros:  88,
					LowMicros:   70,
					CloseMicros: 72,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 9, 10, 0, 0, time.UTC),
					Volume:      85,
					OpenMicros:  98,
					HighMicros:  123,
					LowMicros:   98,
					CloseMicros: 98,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 15, 10, 0, 0, time.UTC),
					Volume:      775,
					OpenMicros:  951,
					HighMicros:  951,
					LowMicros:   515,
					CloseMicros: 723,
				},
				{
					OpenedAt:    time.Date(2020, 7, 8, 10, 10, 0, 0, time.UTC),
					Volume:      975,
					OpenMicros:  289,
					HighMicros:  321,
					LowMicros:   77,
					CloseMicros: 289,
				},
			},
		),
	)
	t.Run(
		"multiple candles for multiple days provided (10-minute)",
		testCompressCandlesFunc(
			testCandles,
			10,
			"minute",
			[]Candle{
				{
					OpenedAt:    time.Date(2020, 7, 1, 9, 0, 0, 0, time.UTC),
					Volume:      50,
					OpenMicros:  11,
					HighMicros:  34,
					LowMicros:   7,
					CloseMicros: 33,
				},
				{
					OpenedAt: time.Date(2020, 7, 1, 10, 0, 0, 0, time.UTC),
					// 510 + 223 = 733
					Volume:      733,
					OpenMicros:  344,
					HighMicros:  588,
					LowMicros:   77,
					CloseMicros: 79,
				},
				{
					OpenedAt:    time.Date(2020, 7, 1, 16, 30, 0, 0, time.UTC),
					Volume:      112,
					OpenMicros:  34,
					HighMicros:  88,
					LowMicros:   70,
					CloseMicros: 72,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 9, 10, 0, 0, time.UTC),
					Volume:      85,
					OpenMicros:  98,
					HighMicros:  123,
					LowMicros:   98,
					CloseMicros: 98,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 15, 10, 0, 0, time.UTC),
					Volume:      775,
					OpenMicros:  951,
					HighMicros:  951,
					LowMicros:   515,
					CloseMicros: 723,
				},
				{
					OpenedAt:    time.Date(2020, 7, 8, 10, 10, 0, 0, time.UTC),
					Volume:      975,
					OpenMicros:  289,
					HighMicros:  321,
					LowMicros:   77,
					CloseMicros: 289,
				},
			},
		),
	)
	t.Run(
		"multiple candles for multiple days provided (15-minute)",
		testCompressCandlesFunc(
			testCandles,
			15,
			"minute",
			[]Candle{
				{
					OpenedAt:    time.Date(2020, 7, 1, 9, 0, 0, 0, time.UTC),
					Volume:      50,
					OpenMicros:  11,
					HighMicros:  34,
					LowMicros:   7,
					CloseMicros: 33,
				},
				{
					OpenedAt: time.Date(2020, 7, 1, 10, 0, 0, 0, time.UTC),
					// 510 + 223 = 733
					Volume:      733,
					OpenMicros:  344,
					HighMicros:  588,
					LowMicros:   77,
					CloseMicros: 79,
				},
				{
					OpenedAt:    time.Date(2020, 7, 1, 16, 30, 0, 0, time.UTC),
					Volume:      112,
					OpenMicros:  34,
					HighMicros:  88,
					LowMicros:   70,
					CloseMicros: 72,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 9, 0, 0, 0, time.UTC),
					Volume:      85,
					OpenMicros:  98,
					HighMicros:  123,
					LowMicros:   98,
					CloseMicros: 98,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 15, 0, 0, 0, time.UTC),
					Volume:      775,
					OpenMicros:  951,
					HighMicros:  951,
					LowMicros:   515,
					CloseMicros: 723,
				},
				{
					OpenedAt:    time.Date(2020, 7, 8, 10, 0, 0, 0, time.UTC),
					Volume:      975,
					OpenMicros:  289,
					HighMicros:  321,
					LowMicros:   77,
					CloseMicros: 289,
				},
			},
		),
	)
	t.Run(
		"multiple candles for multiple days provided (30-minute)",
		testCompressCandlesFunc(
			testCandles,
			30,
			"minute",
			[]Candle{
				{
					OpenedAt:    time.Date(2020, 7, 1, 9, 0, 0, 0, time.UTC),
					Volume:      50,
					OpenMicros:  11,
					HighMicros:  34,
					LowMicros:   7,
					CloseMicros: 33,
				},
				{
					OpenedAt: time.Date(2020, 7, 1, 10, 0, 0, 0, time.UTC),
					// 510 + 223 = 733
					Volume:      733,
					OpenMicros:  344,
					HighMicros:  588,
					LowMicros:   77,
					CloseMicros: 79,
				},
				{
					OpenedAt:    time.Date(2020, 7, 1, 16, 30, 0, 0, time.UTC),
					Volume:      112,
					OpenMicros:  34,
					HighMicros:  88,
					LowMicros:   70,
					CloseMicros: 72,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 9, 0, 0, 0, time.UTC),
					Volume:      85,
					OpenMicros:  98,
					HighMicros:  123,
					LowMicros:   98,
					CloseMicros: 98,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 15, 0, 0, 0, time.UTC),
					Volume:      775,
					OpenMicros:  951,
					HighMicros:  951,
					LowMicros:   515,
					CloseMicros: 723,
				},
				{
					OpenedAt:    time.Date(2020, 7, 8, 10, 0, 0, 0, time.UTC),
					Volume:      975,
					OpenMicros:  289,
					HighMicros:  321,
					LowMicros:   77,
					CloseMicros: 289,
				},
			},
		),
	)
	t.Run(
		"multiple candles for multiple days provided (1-hour)",
		testCompressCandlesFunc(
			testCandles,
			1,
			"hour",
			[]Candle{
				{
					OpenedAt:    time.Date(2020, 7, 1, 9, 0, 0, 0, time.UTC),
					Volume:      50,
					OpenMicros:  11,
					HighMicros:  34,
					LowMicros:   7,
					CloseMicros: 33,
				},
				{
					OpenedAt: time.Date(2020, 7, 1, 10, 0, 0, 0, time.UTC),
					// 510 + 223 = 733
					Volume:      733,
					OpenMicros:  344,
					HighMicros:  588,
					LowMicros:   77,
					CloseMicros: 79,
				},
				{
					OpenedAt:    time.Date(2020, 7, 1, 16, 0, 0, 0, time.UTC),
					Volume:      112,
					OpenMicros:  34,
					HighMicros:  88,
					LowMicros:   70,
					CloseMicros: 72,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 9, 0, 0, 0, time.UTC),
					Volume:      85,
					OpenMicros:  98,
					HighMicros:  123,
					LowMicros:   98,
					CloseMicros: 98,
				},
				{
					OpenedAt:    time.Date(2020, 7, 2, 15, 0, 0, 0, time.UTC),
					Volume:      775,
					OpenMicros:  951,
					HighMicros:  951,
					LowMicros:   515,
					CloseMicros: 723,
				},
				{
					OpenedAt:    time.Date(2020, 7, 8, 10, 0, 0, 0, time.UTC),
					Volume:      975,
					OpenMicros:  289,
					HighMicros:  321,
					LowMicros:   77,
					CloseMicros: 289,
				},
			},
		),
	)
}

func testCompressCandlesFunc(candles []Candle, timeInterval uint, timeUnit string, expected []Candle) func(*testing.T) {
	return func(t *testing.T) {
		actual, err := CompressCandles(candles, timeInterval, timeUnit, time.UTC)
		if err != nil || !eqCandleSlice(expected, actual) {
			t.Errorf("\nExpected: %+v\nActual: %v", expected, actual)
		}
	}
}

func TestGetMidnight(t *testing.T) {
	midnight := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
	midnightResult := utils.GetMidnight(midnight, time.UTC)
	if midnightResult != midnight {
		t.Errorf("Midnight input failed, expected %s, got %s\n", midnight, midnightResult)
	}
	quarterPastNine := time.Date(2020, 7, 1, 9, 15, 0, 0, time.UTC)
	quarterPastNineResult := utils.GetMidnight(quarterPastNine, time.UTC)
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
			OpenedAt:    time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC),
			Volume:      135,
			OpenMicros:  456,
			HighMicros:  789,
			LowMicros:   753,
			CloseMicros: 159,
		},
		now,
	)
	expectedResult := Candle{
		OpenedAt:    now,
		Volume:      0,
		OpenMicros:  159,
		HighMicros:  159,
		LowMicros:   159,
		CloseMicros: 159,
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
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), LowMicros: 951, Volume: 852},
			},
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), LowMicros: 951, Volume: 852},
			},
		),
	)
	t.Run(
		"Single day with single candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 3), LowMicros: 951, Volume: 852},
				{OpenedAt: addMinutes(now, 5), LowMicros: 888, Volume: 777},
			},
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), OpenMicros: 852, HighMicros: 852, LowMicros: 852, CloseMicros: 852},
				{OpenedAt: addMinutes(now, 3), LowMicros: 951, Volume: 852},
				{OpenedAt: addMinutes(now, 4)},
				{OpenedAt: addMinutes(now, 5), LowMicros: 888, Volume: 777},
			},
		),
	)
	t.Run(
		"Single day with multiple candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 3), LowMicros: 951, Volume: 852},
				{OpenedAt: addMinutes(now, 5), LowMicros: 888, Volume: 777},
			},
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), OpenMicros: 789, HighMicros: 789, LowMicros: 789, CloseMicros: 789},
				{OpenedAt: addMinutes(now, 2), OpenMicros: 789, HighMicros: 789, LowMicros: 789, CloseMicros: 789},
				{OpenedAt: addMinutes(now, 3), LowMicros: 951, Volume: 852},
				{OpenedAt: addMinutes(now, 4)},
				{OpenedAt: addMinutes(now, 5), LowMicros: 888, Volume: 777},
			},
		),
	)
	t.Run(
		"Multiple days with no candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), LowMicros: 951, Volume: 852},
				{OpenedAt: tomorrow, CloseMicros: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), OpenMicros: 788, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 2), LowMicros: 987, Volume: 1},
				{OpenedAt: nextWeek, HighMicros: 999, Volume: 2},
			},
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), LowMicros: 951, Volume: 852},
				{OpenedAt: tomorrow, CloseMicros: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), OpenMicros: 788, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 2), LowMicros: 987, Volume: 1},
				{OpenedAt: nextWeek, HighMicros: 999, Volume: 2},
			},
		),
	)
	t.Run(
		"Multiple days with single candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), LowMicros: 951, Volume: 852},
				{OpenedAt: tomorrow, CloseMicros: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 2), LowMicros: 987, Volume: 1},
				{OpenedAt: nextWeek, HighMicros: 999, Volume: 2},
			},
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), LowMicros: 951, Volume: 852},
				{OpenedAt: tomorrow, CloseMicros: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), OpenMicros: 456, HighMicros: 456, LowMicros: 456, CloseMicros: 456},
				{OpenedAt: addMinutes(tomorrow, 2), LowMicros: 987, Volume: 1},
				{OpenedAt: nextWeek, HighMicros: 999, Volume: 2},
			},
		),
	)
	t.Run(
		"Multiple days with multiple candle gaps",
		testFillMinuteCandlesFunc(
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), LowMicros: 951, Volume: 852},
				{OpenedAt: tomorrow, CloseMicros: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), LowMicros: 777, CloseMicros: 774, Volume: 10},
				{OpenedAt: addMinutes(tomorrow, 5), LowMicros: 987, Volume: 1},
				{OpenedAt: nextWeek, HighMicros: 999, Volume: 2},
				{OpenedAt: addMinutes(nextWeek, 2), HighMicros: 77, Volume: 50},
			},
			[]Candle{
				{OpenedAt: now, OpenMicros: 159, HighMicros: 123, LowMicros: 456, CloseMicros: 789, Volume: 753},
				{OpenedAt: addMinutes(now, 1), CloseMicros: 852, Volume: 357},
				{OpenedAt: addMinutes(now, 2), LowMicros: 951, Volume: 852},
				{OpenedAt: tomorrow, CloseMicros: 456, Volume: 325},
				{OpenedAt: addMinutes(tomorrow, 1), LowMicros: 777, CloseMicros: 774, Volume: 10},
				{OpenedAt: addMinutes(tomorrow, 2), OpenMicros: 774, HighMicros: 774, LowMicros: 774, CloseMicros: 774},
				{OpenedAt: addMinutes(tomorrow, 3), OpenMicros: 774, HighMicros: 774, LowMicros: 774, CloseMicros: 774},
				{OpenedAt: addMinutes(tomorrow, 4), OpenMicros: 774, HighMicros: 774, LowMicros: 774, CloseMicros: 774},
				{OpenedAt: addMinutes(tomorrow, 5), LowMicros: 987, Volume: 1},
				{OpenedAt: nextWeek, HighMicros: 999, Volume: 2},
				{OpenedAt: addMinutes(nextWeek, 1)},
				{OpenedAt: addMinutes(nextWeek, 2), HighMicros: 77, Volume: 50},
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

func addMinutes(moment time.Time, minutes int) time.Time {
	return moment.Add(time.Duration(minutes) * time.Minute)
}

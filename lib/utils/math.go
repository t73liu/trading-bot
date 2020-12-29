package utils

import "math"

const Million = 1000000

func RoundToTwoDecimals(value float64) float64 {
	return RoundToNDecimals(value, 2)
}

func RoundToNDecimals(value float64, decimalPlaces int) float64 {
	offset := math.Pow10(decimalPlaces)
	return math.Round(value*offset) / offset
}

func DollarsToMicros(dollars float64) int64 {
	return int64(dollars * Million)
}

func MicrosToDollars(micros int64) float64 {
	return float64(micros) / Million
}

func Sum(nums ...int64) (total int64) {
	for _, num := range nums {
		total += num
	}
	return total
}

func Abs(num int64) int64 {
	if num < 0 {
		return -1 * num
	}
	return num
}

func Max(nums ...int64) int64 {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

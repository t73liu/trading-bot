package utils

import "math"

const Million = 1000000

func RoundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
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

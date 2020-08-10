package analyze

func genValidMicro(value int64) ValidMicro {
	return ValidMicro{
		Micro: value,
		Valid: true,
	}
}

func eqValidCalcSlice(expected, actual []ValidMicro) bool {
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

func genValidFloat(value float64) ValidFloat {
	return ValidFloat{
		Value: value,
		Valid: true,
	}
}

func eqValidFloatSlice(expected, actual []ValidFloat) bool {
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

func calcSum(values []int64, startIndex, endIndex int) (sum int64) {
	for i := startIndex; i < len(values) && i <= endIndex; i++ {
		sum += values[i]
	}
	return sum
}

func calcAverage(values []int64, startIndex, endIndex int) (sum int64) {
	return calcSum(values, startIndex, endIndex) / int64(endIndex-startIndex+1)
}

func maxInt(nums ...int64) int64 {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

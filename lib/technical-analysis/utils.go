package analyze

func eqMicroDollarRangeSlice(expected, actual []MicroDollarRange) bool {
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

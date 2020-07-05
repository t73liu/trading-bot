package analyze

func genValidCalc(value int64) ValidCalc {
	return ValidCalc{
		Calc:  value,
		Valid: true,
	}
}

func eqValidCalcSlice(expected, actual []ValidCalc) bool {
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

const million = 1000000

func dollarsToMicros(dollars float64) int64 {
	return int64(dollars * million)
}

func microsToDollars(micros int64) float64 {
	return float64(micros / million)
}

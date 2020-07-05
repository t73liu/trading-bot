package analyze

type Candle struct {
	Volume int
	Open   int64
	High   int64
	Low    int64
	Close  int64
}

type ValidCalc struct {
	Calc  int64
	Valid bool
}

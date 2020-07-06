## Technical Analysis

Go module containing common technical analysis indicators.

Note: `int64` are used to represent dollar micros (i.e. 1,000,000 is equivalent to 1.00)

Note: If there are gaps in the candlestick data, then no change in the price
is assumed (i.e. same price for open, high, low, close and volume is zero).

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/technical-analysis"

func main() {
	candles := []analyze.Candle{{Close: 12}, {Close: 15}}
	// SMA 20-period
	results := analyze.SMA(candles, 20)
	// EMA 20-period
	results := analyze.EMA(candles, 20)
}
```

## Indicators

- [x] Simple Moving Average (SMA)
- [x] Exponential Moving Average (EMA)
- [x] Relative Strength Index (RSI)
- [x] Moving Average Convergence Divergence (MACD)
- [ ] Stochastic Oscillator
- [ ] Volume Weighted Average Price (VWAP)
- [ ] Commodity Channel Index (CCI)

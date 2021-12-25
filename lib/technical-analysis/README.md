# Technical Analysis

Go module containing common technical analysis indicators.

Note: `int64` are used to represent dollar micros (i.e. 1,000,000 is equivalent to 1.00)

Note: If there are gaps in the candlestick data, then no change in the price
is assumed (i.e. same price for open, high, low, close and volume is zero).

## Usage

```golang
package main

import "github.com/t73liu/trading-bot/lib/candle"
import "github.com/t73liu/trading-bot/lib/technical-analysis"

func main() {
	candles := []candle.Candle{{Close: 12}, {Close: 15}}
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
- [x] Volume Weighted Average Price (VWAP)
- [x] Average True Value (ATR)
- [x] Keltner Channels
- [x] Bollinger Bands
- [x] TTM Squeeze
- [ ] Ease of Movement (EMV)
- [ ] Stochastic Oscillator
- [ ] Commodity Channel Index (CCI)
- [ ] Ehler's Roofing Filter

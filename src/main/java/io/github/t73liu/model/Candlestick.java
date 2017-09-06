package io.github.t73liu.model;

import java.math.BigDecimal;
import java.math.RoundingMode;

public class Candlestick {
    private final BigDecimal open;
    private final BigDecimal close;
    private final BigDecimal high;
    private final BigDecimal low;
    private CandlestickType type;

    public Candlestick(double open, double close, double high, double low) {
        this.open = new BigDecimal(open).setScale(8, RoundingMode.HALF_UP);
        this.close = new BigDecimal(close).setScale(8, RoundingMode.HALF_UP);
        this.high = new BigDecimal(high).setScale(8, RoundingMode.HALF_UP);
        this.low = new BigDecimal(low).setScale(8, RoundingMode.HALF_UP);
    }

    public BigDecimal getOpen() {
        return open;
    }

    public BigDecimal getClose() {
        return close;
    }

    public BigDecimal getHigh() {
        return high;
    }

    public BigDecimal getLow() {
        return low;
    }

    public CandlestickType getType() {
        return type;
    }

    public void setType(CandlestickType type) {
        this.type = type;
    }
}

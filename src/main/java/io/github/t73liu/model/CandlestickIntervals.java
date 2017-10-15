package io.github.t73liu.model;

public enum CandlestickIntervals {
    FIVE_MIN(300),
    FIFTEEN_MIN(900),
    THIRTY_MIN(1800),
    TWO_HOUR(7200); // NOT IN BITFINEX

    private final long interval;

    CandlestickIntervals(long interval) {
        this.interval = interval;
    }

    public long getInterval() {
        return interval;
    }
}

package io.github.t73liu.model;

public enum CandlestickIntervals {
    FIVE_MIN(300),
    FIFTEEN_MIN(900),
    THIRTY_MIN(1800),
    //    ONE_HOUR("3600"), // Not supported by poloniex api
    TWO_HOUR(7200);

    private final long interval;

    CandlestickIntervals(long interval) {
        this.interval = interval;
    }

    public long getInterval() {
        return interval;
    }
}

package io.github.t73liu.model.bitfinex;

public enum BitfinexCandleInterval {
    ONE_MIN("1m"),
    FIVE_MIN("5m"),
    FIFTEEN_MIN("15m"),
    THIRTY_MIN("30m"),
    ONE_HOUR("1h"),
    THREE_HOUR("3h"),
    SIX_HOUR("6h"),
    TWELVE_HOUR("12h"),
    ONE_DAY("1D"),
    SEVEN_DAY("7D"),
    TWO_WEEK("14D"),
    ONE_MONTH("1M");

    private final String intervalName;

    BitfinexCandleInterval(String intervalName) {
        this.intervalName = intervalName;
    }

    public String getIntervalName() {
        return intervalName;
    }
}

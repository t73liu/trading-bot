package io.github.t73liu.model.poloniex;

public enum PoloniexCandleInterval {
    FIVE_MIN("300"),
    FIFTEEN_MIN("900"),
    THIRTY_MIN("1800"),
    TWO_HOUR("7200"),
    FOUR_HOUR("14400"),
    ONE_DAY("86400");

    private final String intervalName;

    PoloniexCandleInterval(String intervalName) {
        this.intervalName = intervalName;
    }

    public String getIntervalName() {
        return intervalName;
    }
}

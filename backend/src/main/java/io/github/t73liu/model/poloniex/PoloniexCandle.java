package io.github.t73liu.model.poloniex;

import io.github.t73liu.util.DateUtil;
import org.ta4j.core.Bar;
import org.ta4j.core.BaseBar;
import org.ta4j.core.num.DoubleNum;

public class PoloniexCandle {
    private long date;
    private double high;
    private double low;
    private double open;
    private double close;
    private double volume;
    private double quote;
    private double weightedAverage;

    public long getDate() {
        return date;
    }

    public double getHigh() {
        return high;
    }

    public double getLow() {
        return low;
    }

    public double getOpen() {
        return open;
    }

    public double getClose() {
        return close;
    }

    public double getVolume() {
        return volume;
    }

    public double getQuote() {
        return quote;
    }

    public double getWeightedAverage() {
        return weightedAverage;
    }

    public Bar toTick() {
        return new BaseBar(DateUtil.unixSecondsToZonedDateTime(date), open, high, low, close, volume, DoubleNum::valueOf);
    }
}

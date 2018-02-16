package io.github.t73liu.model.bitfinex;

import org.ta4j.core.Bar;
import org.ta4j.core.BaseBar;
import org.ta4j.core.num.DoubleNum;

import static io.github.t73liu.util.DateUtil.unixMillisecondsToZonedDateTime;

public class BitfinexCandle {
    private long millisecondTimeStamp;
    private double open;
    private double close;
    private double high;
    private double low;
    private double volume;

    public BitfinexCandle(double[] jsonArray) {
        this.millisecondTimeStamp = (long) jsonArray[0];
        this.open = jsonArray[1];
        this.close = jsonArray[2];
        this.high = jsonArray[3];
        this.low = jsonArray[4];
        this.volume = jsonArray[5];
    }

    public long getMillisecondTimeStamp() {
        return millisecondTimeStamp;
    }

    public double getOpen() {
        return open;
    }

    public double getClose() {
        return close;
    }

    public double getHigh() {
        return high;
    }

    public double getLow() {
        return low;
    }

    public double getVolume() {
        return volume;
    }

    public Bar toTick() {
        return new BaseBar(unixMillisecondsToZonedDateTime(millisecondTimeStamp), open, high, low, close, volume, DoubleNum::valueOf);
    }
}

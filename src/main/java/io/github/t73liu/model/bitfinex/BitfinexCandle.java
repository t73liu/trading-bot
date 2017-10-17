package io.github.t73liu.model.bitfinex;

import eu.verdelhan.ta4j.BaseTick;
import eu.verdelhan.ta4j.Tick;
import io.github.t73liu.util.DateUtil;

public class BitfinexCandle {
    private long millisecondTimeStamp;
    private double open;
    private double close;
    private double high;
    private double low;
    private double volume;

    public BitfinexCandle(Object[] jsonArray) {
        this.millisecondTimeStamp = (long) jsonArray[0];
        this.open = (double) jsonArray[1];
        this.close = (double) jsonArray[2];
        this.high = (double) jsonArray[3];
        this.low = (double) jsonArray[4];
        this.volume = (double) jsonArray[5];
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

    public Tick toTick() {
        return new BaseTick(DateUtil.unixTimeStampToZonedDateTime(millisecondTimeStamp / 1000), open, high, low, close, volume);
    }
}

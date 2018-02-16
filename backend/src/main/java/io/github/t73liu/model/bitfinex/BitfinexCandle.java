package io.github.t73liu.model.bitfinex;

import com.fasterxml.jackson.annotation.JsonCreator;
import io.github.t73liu.model.Candlestick;

import java.time.ZonedDateTime;

import static io.github.t73liu.util.DateUtil.unixMillisecondsToZonedDateTime;

public class BitfinexCandle implements Candlestick {
    private ZonedDateTime dateTime;
    private double open;
    private double close;
    private double high;
    private double low;
    private double volume;

    @JsonCreator
    public BitfinexCandle(double[] jsonArray) {
        this.dateTime = unixMillisecondsToZonedDateTime((long) jsonArray[0]);
        this.open = jsonArray[1];
        this.close = jsonArray[2];
        this.high = jsonArray[3];
        this.low = jsonArray[4];
        this.volume = jsonArray[5];
    }

    @Override
    public ZonedDateTime getDateTime() {
        return dateTime;
    }

    @Override
    public double getOpen() {
        return open;
    }

    @Override
    public double getHigh() {
        return high;
    }

    @Override
    public double getLow() {
        return low;
    }

    @Override
    public double getClose() {
        return close;
    }

    @Override
    public double getVolume() {
        return volume;
    }
}

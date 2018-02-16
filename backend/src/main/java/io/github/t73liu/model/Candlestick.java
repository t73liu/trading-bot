package io.github.t73liu.model;

import org.ta4j.core.Bar;
import org.ta4j.core.BaseBar;
import org.ta4j.core.num.DoubleNum;

import java.time.ZonedDateTime;

public interface Candlestick {
    static Bar convertToBar(Candlestick candlestick) {
        return new BaseBar(candlestick.getDateTime(), candlestick.getOpen(), candlestick.getHigh(), candlestick.getLow(),
                candlestick.getClose(), candlestick.getVolume(), DoubleNum::valueOf);
    }

    ZonedDateTime getDateTime();

    double getOpen();

    double getHigh();

    double getLow();

    double getClose();

    double getVolume();
}

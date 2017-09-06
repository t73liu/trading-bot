package io.github.t73liu.strategy.candlestick;

import io.github.t73liu.model.Candlestick;

import java.util.List;

import static io.github.t73liu.model.CandlestickType.*;

public class CandlestickProcessor {
    public static List<Candlestick> processCandlesticks(List<Candlestick> candlesticks) {
        for (Candlestick candle : candlesticks) {
            double open = candle.getOpen().doubleValue();
            double close = candle.getClose().doubleValue();
            double low = candle.getLow().doubleValue();
            double high = candle.getHigh().doubleValue();
            double rectangle = close - open;
            double upperWick = high - (rectangle > 0 ? close : open);
            double lowerWick = Math.abs(low - (rectangle < 0 ? close : open));
            double center = (high + low) / 2;
            if (rectangle > 0) {
                candle.setType(BULL);
            } else if (rectangle < 0) {
                candle.setType(BEAR);
            } else {
                candle.setType(NEUTRAL);
            }
        }
        return candlesticks;
    }
}

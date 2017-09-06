package io.github.t73liu.strategy.candlestick;

import io.github.t73liu.model.Candlestick;

import java.util.List;

import static io.github.t73liu.model.CandlestickType.BUY;

public class CandlestickProcessor {
    public static List<Candlestick> processCandlesticks(List<Candlestick> candlesticks) {
        for (Candlestick candle : candlesticks) {
            candle.setType(BUY);
        }
        return candlesticks;
    }
}

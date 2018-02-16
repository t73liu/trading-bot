package io.github.t73liu.model.poloniex;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import io.github.t73liu.model.Candlestick;

import java.time.ZonedDateTime;

import static io.github.t73liu.util.DateUtil.unixSecondsToZonedDateTime;

public class PoloniexCandle implements Candlestick {
    private ZonedDateTime dateTime;
    private double high;
    private double low;
    private double open;
    private double close;
    private double volume;
    private double quote;
    private double weightedAverage;

    @JsonCreator
    public PoloniexCandle(@JsonProperty("date") long date, @JsonProperty("high") double high, @JsonProperty("low") double low,
                          @JsonProperty("open") double open, @JsonProperty("close") double close, @JsonProperty("volume") double volume,
                          @JsonProperty("quoteVolume") double quote, @JsonProperty("weightedAverage") double weightedAverage) {
        this.dateTime = unixSecondsToZonedDateTime(date);
        this.high = high;
        this.low = low;
        this.open = open;
        this.close = close;
        this.volume = volume;
        this.quote = quote;
        this.weightedAverage = weightedAverage;
    }

    @Override
    public ZonedDateTime getDateTime() {
        return dateTime;
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
    public double getOpen() {
        return open;
    }

    @Override
    public double getClose() {
        return close;
    }

    @Override
    public double getVolume() {
        return volume;
    }

    public double getQuote() {
        return quote;
    }

    public double getWeightedAverage() {
        return weightedAverage;
    }
}

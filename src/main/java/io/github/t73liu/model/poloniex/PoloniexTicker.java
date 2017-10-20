package io.github.t73liu.model.poloniex;

public class PoloniexTicker {
    private double last;
    private double lowestAsk;
    private double highestBid;
    private double percentageChange;
    private double baseVolume;
    private double quoteVolume;

    public double getLast() {
        return last;
    }

    public double getLowestAsk() {
        return lowestAsk;
    }

    public double getHighestBid() {
        return highestBid;
    }

    public double getPercentageChange() {
        return percentageChange;
    }

    public double getBaseVolume() {
        return baseVolume;
    }

    public double getQuoteVolume() {
        return quoteVolume;
    }
}

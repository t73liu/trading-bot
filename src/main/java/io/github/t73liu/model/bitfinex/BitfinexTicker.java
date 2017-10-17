package io.github.t73liu.model.bitfinex;

public class BitfinexTicker {
    private String symbol;
    private double bid;
    private double bidSize;
    private double ask;
    private double askSize;
    private double dailyChange;
    private double dailyChangePercent;
    private double lastPrice;
    private double volume;
    private double high;
    private double low;

    public BitfinexTicker(Object[] jsonArray) {
        this.symbol = jsonArray[0].toString();
        this.bid = (double) jsonArray[1];
        this.bidSize = (double) jsonArray[2];
        this.ask = (double) jsonArray[3];
        this.askSize = (double) jsonArray[4];
        this.dailyChange = (double) jsonArray[5];
        this.dailyChangePercent = (double) jsonArray[6];
        this.lastPrice = (double) jsonArray[7];
        this.volume = (double) jsonArray[8];
        this.high = (double) jsonArray[9];
        this.low = (double) jsonArray[10];
    }

    public String getSymbol() {
        return symbol;
    }

    public double getBid() {
        return bid;
    }

    public double getBidSize() {
        return bidSize;
    }

    public double getAsk() {
        return ask;
    }

    public double getAskSize() {
        return askSize;
    }

    public double getDailyChange() {
        return dailyChange;
    }

    public double getDailyChangePercent() {
        return dailyChangePercent;
    }

    public double getLastPrice() {
        return lastPrice;
    }

    public double getVolume() {
        return volume;
    }

    public double getHigh() {
        return high;
    }

    public double getLow() {
        return low;
    }
}

package io.github.t73liu.model.bitfinex;

public class BitfinexTicker {
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

    public BitfinexTicker(double[] jsonArray) {
        this.bid = jsonArray[0];
        this.bidSize = jsonArray[1];
        this.ask = jsonArray[2];
        this.askSize = jsonArray[3];
        this.dailyChange = jsonArray[4];
        this.dailyChangePercent = jsonArray[5];
        this.lastPrice = jsonArray[6];
        this.volume = jsonArray[7];
        this.high = jsonArray[8];
        this.low = jsonArray[9];
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

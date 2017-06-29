package io.github.t73liu.model;

// Might be exchange specific
public class CurrencyAnalytics {
    private double high;
    private double low;
    private double dailyVolume;

    public double getHigh() {
        return high;
    }

    public void setHigh(double high) {
        this.high = high;
    }

    public double getLow() {
        return low;
    }

    public void setLow(double low) {
        this.low = low;
    }

    public double getDailyVolume() {
        return dailyVolume;
    }

    public void setDailyVolume(double dailyVolume) {
        this.dailyVolume = dailyVolume;
    }
}

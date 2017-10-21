package io.github.t73liu.model.quadriga;

public class QuadrigaTicker {
    private long timestamp;
    private double high;
    private double last;
    private double low;
    private double vwap;
    private double volume;
    private double bid;
    private double ask;

    public long getTimestamp() {
        return timestamp;
    }

    public double getHigh() {
        return high;
    }

    public double getLast() {
        return last;
    }

    public double getLow() {
        return low;
    }

    public double getVwap() {
        return vwap;
    }

    public double getVolume() {
        return volume;
    }

    public double getBid() {
        return bid;
    }

    public double getAsk() {
        return ask;
    }
}

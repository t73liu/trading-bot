package io.github.t73liu.model.quadriga;

import java.util.List;

public class QuadrigaOrderBook {
    private long timestamp;
    private List<double[]> asks;
    private List<double[]> bids;

    public long getTimestamp() {
        return timestamp;
    }

    public List<double[]> getAsks() {
        return asks;
    }

    public List<double[]> getBids() {
        return bids;
    }
}

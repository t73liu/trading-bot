package io.github.t73liu.model.poloniex;

import java.util.List;

public class PoloniexOrderBook {
    private List<double[]> asks;
    private List<double[]> bids;
    private int isFrozen;
    private int seq;

    public List<double[]> getAsks() {
        return asks;
    }

    public List<double[]> getBids() {
        return bids;
    }

    public int getIsFrozen() {
        return isFrozen;
    }

    public int getSeq() {
        return seq;
    }
}

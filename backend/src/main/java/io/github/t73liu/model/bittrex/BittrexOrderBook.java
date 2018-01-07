package io.github.t73liu.model.bittrex;

import java.util.List;

public class BittrexOrderBook {
    private boolean success;
    private String message;
    private BittrexOrderBookResult result;

    public boolean isSuccess() {
        return success;
    }

    public String getMessage() {
        return message;
    }

    public BittrexOrderBookResult getResult() {
        return result;
    }

    public static class BittrexOrderBookResult {
        private List<BittrexOrderBookEntry> buy;
        private List<BittrexOrderBookEntry> sell;

        public List<BittrexOrderBookEntry> getBuy() {
            return buy;
        }

        public List<BittrexOrderBookEntry> getSell() {
            return sell;
        }
    }

    public static class BittrexOrderBookEntry {
        private double quantity;
        private double rate;

        public double getQuantity() {
            return quantity;
        }

        public double getRate() {
            return rate;
        }
    }
}

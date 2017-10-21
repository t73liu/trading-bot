package io.github.t73liu.model.bittrex;

public class BittrexTicker {
    private boolean success;
    private String message;
    private BittrexTickerResult result;

    public boolean isSuccess() {
        return success;
    }

    public String getMessage() {
        return message;
    }

    public BittrexTickerResult getResult() {
        return result;
    }

    public static class BittrexTickerResult {
        private double bid;
        private double ask;
        private double last;

        public double getBid() {
            return bid;
        }

        public double getAsk() {
            return ask;
        }

        public double getLast() {
            return last;
        }
    }
}

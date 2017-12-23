package io.github.t73liu.model.bitfinex;

public class BitfinexOrderBook {
    private double price;
    private int count;
    private double amount;
    private String type;

    public BitfinexOrderBook(double[] jsonArray) {
        this.price = jsonArray[0];
        this.count = (int) jsonArray[1];
        this.amount = jsonArray[2];
        this.type = amount > 0 ? "bid" : "ask";
    }

    public double getPrice() {
        return price;
    }

    public int getCount() {
        return count;
    }

    public double getAmount() {
        return amount;
    }

    public String getType() {
        return type;
    }
}

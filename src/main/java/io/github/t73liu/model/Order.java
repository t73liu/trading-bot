package io.github.t73liu.model;

public class Order {
    private String currentCoin;
    private String targetCoin;
    private Double quantity;
    private Double exchangeRate;
    private Exchange exchange;

    public String getCurrentCoin() {
        return currentCoin;
    }

    public void setCurrentCoin(String currentCoin) {
        this.currentCoin = currentCoin;
    }

    public String getTargetCoin() {
        return targetCoin;
    }

    public void setTargetCoin(String targetCoin) {
        this.targetCoin = targetCoin;
    }

    public Double getQuantity() {
        return quantity;
    }

    public void setQuantity(Double quantity) {
        this.quantity = quantity;
    }

    public Double getExchangeRate() {
        return exchangeRate;
    }

    public void setExchangeRate(Double exchangeRate) {
        this.exchangeRate = exchangeRate;
    }

    public Exchange getExchange() {
        return exchange;
    }

    public void setExchange(Exchange exchange) {
        this.exchange = exchange;
    }
}

package io.github.t73liu.model;

import org.apache.commons.lang3.tuple.Pair;

import java.time.LocalDateTime;

public class Order {
    // <Starting Currency, Ending Currency>
    private Pair<String, String> currencyPair;
    private double quantity;
    private double price;
    private LocalDateTime issueTime;
    // Buy vs Sell
    private String type;
    private Exchange exchange;

    public Pair<String, String> getCurrencyPair() {
        return currencyPair;
    }

    public void setCurrencyPair(Pair<String, String> currencyPair) {
        this.currencyPair = currencyPair;
    }

    public double getQuantity() {
        return quantity;
    }

    public void setQuantity(double quantity) {
        this.quantity = quantity;
    }

    public double getPrice() {
        return price;
    }

    public void setPrice(double price) {
        this.price = price;
    }

    public LocalDateTime getIssueTime() {
        return issueTime;
    }

    public void setIssueTime(LocalDateTime issueTime) {
        this.issueTime = issueTime;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Exchange getExchange() {
        return exchange;
    }

    public void setExchange(Exchange exchange) {
        this.exchange = exchange;
    }
}

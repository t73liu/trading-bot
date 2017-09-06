package io.github.t73liu.model;

import java.math.BigDecimal;

public class Balance {
    private String currency;
    private BigDecimal available;
    private BigDecimal onOrders;
    private BigDecimal usdValue;

    public String getCurrency() {
        return currency;
    }

    public void setCurrency(String currency) {
        this.currency = currency;
    }

    public BigDecimal getAvailable() {
        return available;
    }

    public void setAvailable(BigDecimal available) {
        this.available = available;
    }

    public BigDecimal getOnOrders() {
        return onOrders;
    }

    public void setOnOrders(BigDecimal onOrders) {
        this.onOrders = onOrders;
    }

    public BigDecimal getUsdValue() {
        return usdValue;
    }

    public void setUsdValue(BigDecimal usdValue) {
        this.usdValue = usdValue;
    }

    @Override
    public String toString() {
        return "Balance{" +
                "currency='" + currency + '\'' +
                ", available=" + available +
                ", onOrders=" + onOrders +
                ", usdValue=" + usdValue +
                '}';
    }
}

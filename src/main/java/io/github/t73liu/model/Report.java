package io.github.t73liu.model;

import java.time.LocalDateTime;
import java.util.Map;

public class Report {
    private LocalDateTime lastChecked;
    private Map<String, Double> currencyHoldings;
    private double marketValue;

    public LocalDateTime getLastChecked() {
        return lastChecked;
    }

    public void setLastChecked(LocalDateTime lastChecked) {
        this.lastChecked = lastChecked;
    }

    public Map<String, Double> getCurrencyHoldings() {
        return currencyHoldings;
    }

    public void setCurrencyHoldings(Map<String, Double> currencyHoldings) {
        this.currencyHoldings = currencyHoldings;
    }

    public double getMarketValue() {
        return marketValue;
    }

    public void setMarketValue(double marketValue) {
        this.marketValue = marketValue;
    }
}

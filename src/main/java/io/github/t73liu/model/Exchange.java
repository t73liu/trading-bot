package io.github.t73liu.model;

import org.apache.commons.lang3.tuple.Pair;

import java.util.HashSet;
import java.util.Set;

public class Exchange {
    private String name;
    private Fee sellFee;
    private Fee buyFee;
    private Fee cashoutFee;
    private Fee transferFee;
    private Set<Pair<String, String>> supportedCurrencies;

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Fee getSellFee() {
        return sellFee;
    }

    public void setSellFee(Fee sellFee) {
        this.sellFee = sellFee;
    }

    public Fee getBuyFee() {
        return buyFee;
    }

    public void setBuyFee(Fee buyFee) {
        this.buyFee = buyFee;
    }

    public Fee getCashoutFee() {
        return cashoutFee;
    }

    public void setCashoutFee(Fee cashoutFee) {
        this.cashoutFee = cashoutFee;
    }

    public Fee getTransferFee() {
        return transferFee;
    }

    public void setTransferFee(Fee transferFee) {
        this.transferFee = transferFee;
    }

    public Set<Pair<String, String>> getSupportedCurrencies() {
        return supportedCurrencies;
    }

    public void setSupportedCurrencies(Set<Pair<String, String>> supportedCurrencies) {
        this.supportedCurrencies = supportedCurrencies;
    }
}

package io.github.t73liu.model;

public class Order {
    private String currentCoin;
    private String targetCoin;
    private Double fxRate;
    private Fee transactionFee;
    private Fee cashoutFee;

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

    public Double getFxRate() {
        return fxRate;
    }

    public void setFxRate(Double fxRate) {
        this.fxRate = fxRate;
    }

    public Fee getTransactionFee() {
        return transactionFee;
    }

    public void setTransactionFee(Fee transactionFee) {
        this.transactionFee = transactionFee;
    }

    public Fee getCashoutFee() {
        return cashoutFee;
    }

    public void setCashoutFee(Fee cashoutFee) {
        this.cashoutFee = cashoutFee;
    }
}

package io.github.t73liu.model;

public class Fee {
    // Percentage or Fixed or ...
    private String type;
    private double cost;

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public double getCost() {
        return cost;
    }

    public void setCost(double cost) {
        this.cost = cost;
    }
}

package io.github.t73liu.model;

public class RateLimit {
    private int maxRequestsPerHour;

    public int getMaxRequestsPerHour() {
        return maxRequestsPerHour;
    }

    public void setMaxRequestsPerHour(int maxRequestsPerHour) {
        this.maxRequestsPerHour = maxRequestsPerHour;
    }
}

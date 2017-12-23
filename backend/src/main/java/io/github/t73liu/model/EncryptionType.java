package io.github.t73liu.model;

public enum EncryptionType {
    HMAC_SHA256("HmacSHA256"), // Quadriga
    HMAC_SHA384("HmacSHA384"), // Bitfinex
    HMAC_SHA512("HmacSHA512"); // Poloniex, Bittrex

    private final String name;

    EncryptionType(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }
}

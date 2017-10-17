package io.github.t73liu.model;

public enum EncryptionType {
    HMAC_SHA512("HmacSHA512");

    private final String name;

    EncryptionType(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }
}

package io.github.t73liu.service;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "quadrigacx")
public class QuadrigaTicker {
    private static final double FIAT_FEE = 0.005;
    private static final double CRYPTO_FEE = 0.002;
}

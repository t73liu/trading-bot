package io.github.t73liu.service;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "coinbase")
public class CoinbaseService extends ExchangeService {
}

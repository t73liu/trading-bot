package io.github.t73liu.service;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "bittrex")
public class BittrexService extends ExchangeService {
}

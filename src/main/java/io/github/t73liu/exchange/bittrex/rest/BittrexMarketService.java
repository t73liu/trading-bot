package io.github.t73liu.exchange.bittrex.rest;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.exchange.MarketService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "bittrex")
public class BittrexMarketService extends ExchangeService implements MarketService {
    private static final Logger LOGGER = LoggerFactory.getLogger(BittrexMarketService.class);
    private static final double FEE = 0.0025;

    // TODO implement
}

package io.github.t73liu.exchange.bitfinex.rest;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.exchange.MarketService;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "bitfinex")
public class BitfinexMarketService extends ExchangeService implements MarketService {
}

package io.github.t73liu.exchange.poloniex.rest;

import io.github.t73liu.exchange.ExchangeService;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "poloniex")
public class PoloniexMarketService extends ExchangeService {
}

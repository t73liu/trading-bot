package io.github.t73liu.exchange.poloniex.rest;

import io.github.t73liu.exchange.PrivateExchangeService;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "poloniex")
public class PoloniexOrderService extends PrivateExchangeService {
}

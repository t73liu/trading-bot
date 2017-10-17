package io.github.t73liu.exchange.bitfinex.rest;

import io.github.t73liu.exchange.OrderService;
import io.github.t73liu.exchange.PrivateExchangeService;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "bitfinex")
public class BitfinexOrderService extends PrivateExchangeService implements OrderService {
}

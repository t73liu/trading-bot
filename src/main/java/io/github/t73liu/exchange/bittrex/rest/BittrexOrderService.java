package io.github.t73liu.exchange.bittrex.rest;

import io.github.t73liu.exchange.OrderService;
import io.github.t73liu.exchange.PrivateExchangeService;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "bittrex")
public class BittrexOrderService extends PrivateExchangeService implements OrderService {
    // TODO implement
}

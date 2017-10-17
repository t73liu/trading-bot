package io.github.t73liu.exchange.quadriga.rest;

import io.github.t73liu.exchange.OrderService;
import io.github.t73liu.exchange.PrivateExchangeService;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "quadrigacx")
public class QuadrigaOrderService extends PrivateExchangeService implements OrderService {
}

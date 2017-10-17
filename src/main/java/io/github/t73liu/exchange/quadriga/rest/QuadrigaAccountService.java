package io.github.t73liu.exchange.quadriga.rest;

import io.github.t73liu.exchange.AccountService;
import io.github.t73liu.exchange.PrivateExchangeService;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "quadrigacx")
public class QuadrigaAccountService extends PrivateExchangeService implements AccountService {
    // TODO implement
}

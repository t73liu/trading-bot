package io.github.t73liu.exchange.bittrex.rest;

import io.github.t73liu.exchange.AccountService;
import io.github.t73liu.exchange.PrivateExchangeService;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "bittrex")
public class BittrexAccountService extends PrivateExchangeService implements AccountService {
}

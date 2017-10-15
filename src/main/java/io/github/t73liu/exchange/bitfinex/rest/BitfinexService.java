package io.github.t73liu.exchange.bitfinex.rest;

import io.github.t73liu.exchange.ExchangeService;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
@ConfigurationProperties(prefix = "bittrex")
public class BitfinexService extends ExchangeService {
    public static final double TAKER_FEE = 0.002;
    public static final double MAKER_FEE = 0.001;
    private static final Logger LOGGER = LoggerFactory.getLogger(BitfinexService.class);

    public Map getTicker(String pair) {
        return new HashMap();
    }
}

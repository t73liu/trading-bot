package io.github.t73liu.exchange.bittrex.rest;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.model.bittrex.BittrexPair;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
@ConfigurationProperties(prefix = "bittrex")
public class BittrexMarketService extends ExchangeService {
    private static final Logger LOGGER = LoggerFactory.getLogger(BittrexMarketService.class);
    private static final double FEE = 0.0025;

    public Map getTicker(BittrexPair pair) {
        LOGGER.info(pair.getPairName());
        return new HashMap();
    }
}

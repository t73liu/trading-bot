package io.github.t73liu.exchanges;

import io.github.t73liu.model.currency.BittrexPair;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
@ConfigurationProperties(prefix = "bittrex")
public class BittrexService extends ExchangeService {
    private static final Logger LOGGER = LoggerFactory.getLogger(BittrexService.class);
    private static final double FEE = 0.0025;

    public Map getTicker(BittrexPair pair) {
        LOGGER.info(pair.getPairName());
        return new HashMap();
    }
}

package io.github.t73liu.exchange.quadriga.rest;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.model.quadriga.QuadrigaPair;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
@ConfigurationProperties(prefix = "quadrigacx")
public class QuadrigaService extends ExchangeService {
    private static final Logger LOGGER = LoggerFactory.getLogger(QuadrigaService.class);
    private static final double FIAT_FEE = 0.005;
    private static final double CRYPTO_FEE = 0.002;

    public Map getTicker(QuadrigaPair pair) {
        LOGGER.info(pair.getPairName());
        return new HashMap();
    }
}

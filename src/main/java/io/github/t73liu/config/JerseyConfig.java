package io.github.t73liu.config;

import io.github.t73liu.rest.CoinbaseResource;
import io.github.t73liu.rest.KrakenResource;
import io.github.t73liu.rest.QuadrigaResource;
import org.glassfish.jersey.server.ResourceConfig;
import org.springframework.stereotype.Component;

@Component
public class JerseyConfig extends ResourceConfig {
    public JerseyConfig() {
        register(CoinbaseResource.class);
        register(KrakenResource.class);
        register(QuadrigaResource.class);
    }
}

package io.github.t73liu.config;

import io.github.t73liu.provider.LocalDateParamProvider;
import io.github.t73liu.provider.ObjectMapperContextResolver;
import org.glassfish.jersey.server.ResourceConfig;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Configuration;

import javax.ws.rs.Path;

@Configuration
public class JerseyConfig extends ResourceConfig {
    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());

    private ApplicationContext context;

    @Autowired
    public JerseyConfig(ApplicationContext context) {
        this.context = context;
        setupResources();
        registerProviders();
    }

    private void setupResources() {
        context.getBeansWithAnnotation(Path.class).forEach((name, resource) -> {
            LOGGER.info("Registering Jersey Resource: {}", name);
            register(resource);
        });
    }

    private void registerProviders() {
        register(ObjectMapperContextResolver.class);
        register(LocalDateParamProvider.class);
    }
}

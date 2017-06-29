package io.github.t73liu.provider;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.github.t73liu.util.ObjectMapperFactory;

import javax.ws.rs.ext.ContextResolver;
import javax.ws.rs.ext.Provider;

@Provider
public class ObjectMapperContextResolver implements ContextResolver<ObjectMapper> {
    private final ObjectMapper mapper;

    public ObjectMapperContextResolver() {
        mapper = ObjectMapperFactory.getNewInstance();
    }

    @Override
    public ObjectMapper getContext(Class<?> type) {
        return mapper;
    }
}

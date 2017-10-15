package io.github.t73liu.provider;

import com.fasterxml.jackson.databind.ObjectMapper;

import javax.ws.rs.ext.ContextResolver;
import javax.ws.rs.ext.Provider;

import static io.github.t73liu.util.MapperUtil.JSON_MAPPER;

@Provider
public class ObjectMapperContextResolver implements ContextResolver<ObjectMapper> {
    @Override
    public ObjectMapper getContext(Class<?> type) {
        return JSON_MAPPER;
    }
}

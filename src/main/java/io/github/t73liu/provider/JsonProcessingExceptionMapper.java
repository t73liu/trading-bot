package io.github.t73liu.provider;

import com.fasterxml.jackson.core.JsonProcessingException;
import io.github.t73liu.model.ExceptionWrapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;
import javax.ws.rs.ext.ExceptionMapper;
import javax.ws.rs.ext.Provider;

@Provider
public class JsonProcessingExceptionMapper implements ExceptionMapper<JsonProcessingException> {
    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());

    @Override
    public Response toResponse(JsonProcessingException exception) {
        Status status = Status.BAD_REQUEST;
        ExceptionWrapper exceptionWrapper = new ExceptionWrapper(status, exception);
        LOGGER.error("Resource Thrown JsonProcessingException. {}", exceptionWrapper);
        return Response.status(status).type(MediaType.APPLICATION_JSON).entity(exceptionWrapper).build();
    }
}

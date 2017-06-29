package io.github.t73liu.provider;

import io.github.t73liu.model.ExceptionWrapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;
import javax.ws.rs.ext.ExceptionMapper;
import javax.ws.rs.ext.Provider;

@Provider
public class GeneralExceptionMapper implements ExceptionMapper<Exception> {
    private final Logger LOGGER = LoggerFactory.getLogger(GeneralExceptionMapper.class);

    @Override
    public Response toResponse(Exception e) {
        Status status = Status.INTERNAL_SERVER_ERROR;
        ExceptionWrapper exception = new ExceptionWrapper(status, e);
        LOGGER.error("Resource Thrown General Exception. Message:{}", e.getMessage());
        return Response.status(status).type(MediaType.APPLICATION_JSON).entity(exception).build();
    }
}

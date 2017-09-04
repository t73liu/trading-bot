package io.github.t73liu.provider;

import io.github.t73liu.model.ExceptionWrapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import javax.validation.ConstraintViolationException;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import javax.ws.rs.core.Response.Status;
import javax.ws.rs.ext.ExceptionMapper;
import javax.ws.rs.ext.Provider;

@Provider
public class ValidationExceptionMapper implements ExceptionMapper<ConstraintViolationException> {
    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());

    @Override
    public Response toResponse(ConstraintViolationException exception) {
        Status status = Status.BAD_REQUEST;
        ExceptionWrapper exceptionWrapper = new ExceptionWrapper(status, exception);
        exceptionWrapper.setCause(exception.getConstraintViolations().toString());
        LOGGER.error("Resource Thrown ConstraintViolationException. {}", exceptionWrapper, exception);
        return Response.status(status).type(MediaType.APPLICATION_JSON).entity(exceptionWrapper).build();
    }
}

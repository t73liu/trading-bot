package io.github.t73liu.model;

import javax.ws.rs.core.Response.Status;
import java.util.Optional;

public class ExceptionWrapper {
    private final int statusCode;
    private final String statusPhrase;
    private final String message;
    private final String cause;

    public ExceptionWrapper(Status status, Exception exception) {
        this.statusCode = status.getStatusCode();
        this.statusPhrase = status.getReasonPhrase();
        this.message = exception.getClass().getSimpleName() + ": " + exception.getMessage();
        Optional<Throwable> cause = Optional.ofNullable(exception.getCause());
        this.cause = cause.isPresent() ? cause.get().getMessage() : "UNKNOWN CAUSE";
    }

    public int getStatusCode() {
        return statusCode;
    }

    public String getStatusPhrase() {
        return statusPhrase;
    }

    public String getMessage() {
        return message;
    }

    public String getCause() {
        return cause;
    }

    @Override
    public String toString() {
        return "ExceptionWrapper{" +
                "statusCode=" + statusCode +
                ", statusPhrase='" + statusPhrase + '\'' +
                ", message='" + message + '\'' +
                ", cause='" + cause + '\'' +
                '}';
    }
}

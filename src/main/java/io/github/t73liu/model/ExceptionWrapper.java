package io.github.t73liu.model;

import javax.ws.rs.core.Response.Status;
import java.util.Optional;

public class ExceptionWrapper {
    private int statusCode;
    private String statusPhrase;
    private String message;
    private String cause;

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

    public void setStatusCode(int statusCode) {
        this.statusCode = statusCode;
    }

    public String getStatusPhrase() {
        return statusPhrase;
    }

    public void setStatusPhrase(String statusPhrase) {
        this.statusPhrase = statusPhrase;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public String getCause() {
        return cause;
    }

    public void setCause(String cause) {
        this.cause = cause;
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

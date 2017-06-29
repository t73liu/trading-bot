package io.github.t73liu.model;

import javax.ws.rs.core.Response.Status;

public class ExceptionWrapper {
    private int statusCode;
    private String error;
    private String message;

    public ExceptionWrapper(Status status, Exception exception) {
        this.statusCode = status.getStatusCode();
        this.error = status.getReasonPhrase();
        this.message = exception.getClass().getSimpleName() + ": " + exception.getMessage();
    }

    public int getStatusCode() {
        return statusCode;
    }

    public void setStatusCode(int statusCode) {
        this.statusCode = statusCode;
    }

    public String getError() {
        return error;
    }

    public void setError(String error) {
        this.error = error;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}

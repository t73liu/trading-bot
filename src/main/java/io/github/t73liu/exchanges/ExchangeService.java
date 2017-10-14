package io.github.t73liu.exchanges;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.github.t73liu.util.ObjectMapperFactory;

import javax.annotation.PostConstruct;

public abstract class ExchangeService {
    private String baseUrl;

    private String apiKey;

    private String secretKey;

    protected ObjectMapper mapper;

    public String getBaseUrl() {
        return baseUrl;
    }

    public void setBaseUrl(String baseUrl) {
        this.baseUrl = baseUrl;
    }

    public String getApiKey() {
        return apiKey;
    }

    public void setApiKey(String apiKey) {
        this.apiKey = apiKey;
    }

    public String getSecretKey() {
        return secretKey;
    }

    public void setSecretKey(String secretKey) {
        this.secretKey = secretKey;
    }

    @PostConstruct
    private void init() {
        this.mapper = ObjectMapperFactory.getNewInstance();
    }
}

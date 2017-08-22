package io.github.t73liu.service;

import org.apache.commons.codec.binary.Hex;
import org.apache.http.NameValuePair;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Optional;

@Service
@ConfigurationProperties(prefix = "poloniex")
public class PoloniexService extends ExchangeService {
    public Map getBalance() throws Exception {
        List<NameValuePair> queryParams = new ArrayList<>();
        queryParams.add(new BasicNameValuePair("command", "returnBalances"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        }
    }

    public Map getOpenOrders() throws Exception {
        List<NameValuePair> queryParams = new ArrayList<>();
        queryParams.add(new BasicNameValuePair("command", "returnOpenOrders"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(new BasicNameValuePair("currencyPair", "all"));
        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        }
    }

    private HttpPost generatePost(List<NameValuePair> queryParams) throws Exception {
        Mac shaMac = Mac.getInstance("HmacSHA512");
        SecretKeySpec keySpec = new SecretKeySpec(getSecretKey().getBytes(StandardCharsets.UTF_8), "HmacSHA512");
        shaMac.init(keySpec);
        Optional<String> queryParamStr = queryParams.stream()
                .map(Object::toString)
                .reduce((queryOne, queryTwo) -> queryOne + "&" + queryTwo);
        if (!queryParamStr.isPresent()) {
            throw new IllegalArgumentException("Unable to generate query params for Poloniex API: " + queryParams);
        }
        String sign = Hex.encodeHexString(shaMac.doFinal(queryParamStr.get().getBytes(StandardCharsets.UTF_8)));

        HttpPost post = new HttpPost(getTradingUrl());
        post.addHeader("Key", getApiKey());
        post.addHeader("Sign", sign);

        post.setEntity(new UrlEncodedFormEntity(queryParams));
        return post;
    }

    private String getTradingUrl() {
        return getBaseUrl() + "/tradingApi";
    }
}

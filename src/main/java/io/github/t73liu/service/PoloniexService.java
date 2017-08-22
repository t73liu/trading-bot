package io.github.t73liu.service;

import io.github.t73liu.util.ObjectMapperFactory;
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

@Service
@ConfigurationProperties(prefix = "poloniex")
public class PoloniexService extends ExchangeService {
    public Map getBalance() throws Exception {
        String nonce = String.valueOf(System.currentTimeMillis());
        String command = "returnBalances";
        HttpPost post = generatePost(command, nonce);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return ObjectMapperFactory.getNewInstance().readValue(response.getEntity().getContent(), Map.class);
        }
    }

    private HttpPost generatePost(String command, String nonce) throws Exception {
        Mac shaMac = Mac.getInstance("HmacSHA512");
        SecretKeySpec keySpec = new SecretKeySpec(getSecretKey().getBytes(StandardCharsets.UTF_8), "HmacSHA512");
        shaMac.init(keySpec);
        String queryParams = "command=" + command + "&nonce=" + nonce;
        final byte[] macData = shaMac.doFinal(queryParams.getBytes(StandardCharsets.UTF_8));
        String sign = Hex.encodeHexString(macData);

        HttpPost post = new HttpPost(getTradingUrl());
        post.addHeader("Key", getApiKey());
        post.addHeader("Sign", sign);

        List<NameValuePair> params = new ArrayList<>();
        params.add(new BasicNameValuePair("command", command));
        params.add(new BasicNameValuePair("nonce", nonce));
        post.setEntity(new UrlEncodedFormEntity(params));
        return post;
    }

    private String getTradingUrl() {
        return getBaseUrl() + "/tradingApi";
    }
}

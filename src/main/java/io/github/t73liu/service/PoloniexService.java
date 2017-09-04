package io.github.t73liu.service;

import org.apache.commons.codec.binary.Hex;
import org.apache.http.NameValuePair;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.client.utils.URIBuilder;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
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
    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());
    // TODO remove hardcode
    private static final double fee = 0.0025;

    public Map getTickers() throws Exception {
        List<NameValuePair> queryParams = new ArrayList<>();
        queryParams.add(new BasicNameValuePair("command", "returnTicker"));
        HttpGet get = generateGet(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        }
    }

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

    public boolean checkArbitrage() throws Exception {
        Map<String, Map<String, String>> tickers = getTickers();
        Double btc = Double.valueOf(tickers.get("USDT_BTC").get("lowestAsk"));
        Double eth = Double.valueOf(tickers.get("USDT_ETH").get("lowestAsk"));
        Double eth2btc = Double.valueOf(tickers.get("BTC_ETH").get("lowestAsk"));
        LOGGER.info("exchange: {}, actual: {}", eth / btc, eth2btc);
        return ((eth2btc / (eth / btc)) - 1) > fee * 3;
    }

    private HttpPost generatePost(List<NameValuePair> queryParams) throws Exception {
        Optional<String> queryParamStr = queryParams.stream()
                .map(Object::toString)
                .reduce((queryOne, queryTwo) -> queryOne + "&" + queryTwo);
        if (!queryParamStr.isPresent()) {
            throw new IllegalArgumentException("Unable to generate query params for Poloniex API: " + queryParams);
        }

        // Generating special headers
        Mac shaMac = Mac.getInstance("HmacSHA512");
        SecretKeySpec keySpec = new SecretKeySpec(getSecretKey().getBytes(StandardCharsets.UTF_8), "HmacSHA512");
        shaMac.init(keySpec);
        String sign = Hex.encodeHexString(shaMac.doFinal(queryParamStr.get().getBytes(StandardCharsets.UTF_8)));
        HttpPost post = new HttpPost(getTradingUrl());
        post.addHeader("Key", getApiKey());
        post.addHeader("Sign", sign);

        post.setEntity(new UrlEncodedFormEntity(queryParams));
        return post;
    }

    private HttpGet generateGet(List<NameValuePair> queryParams) throws Exception {
        return new HttpGet(new URIBuilder(getPublicUrl()).addParameters(queryParams).build());
    }

    private String getTradingUrl() {
        return getBaseUrl() + "tradingApi";
    }

    private String getPublicUrl() {
        return getBaseUrl() + "public";
    }
}

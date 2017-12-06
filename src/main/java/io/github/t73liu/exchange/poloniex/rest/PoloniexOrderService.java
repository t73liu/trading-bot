package io.github.t73liu.exchange.poloniex.rest;

import io.github.t73liu.exchange.OrderService;
import io.github.t73liu.exchange.PrivateExchangeService;
import io.github.t73liu.model.EncryptionType;
import io.github.t73liu.model.poloniex.PoloniexPair;
import io.github.t73liu.util.HttpUtil;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import static io.github.t73liu.util.MapperUtil.JSON_READER;

@Service
@ConfigurationProperties(prefix = "poloniex")
public class PoloniexOrderService extends PrivateExchangeService implements OrderService {
    public Map getOpenOrders(PoloniexPair pair) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "returnOpenOrders"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(pair == null ? new BasicNameValuePair("currencyPair", "all") : new BasicNameValuePair("currencyPair", pair.getPairName()));

        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            post.releaseConnection();
        }
    }

    private Map cancelOrder(String orderNumber) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "cancelOrder"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(new BasicNameValuePair("orderNumber", orderNumber));

        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            post.releaseConnection();
        }
    }

    private Map moveOrder(String orderNumber, double rate, Optional<Double> amount, String fulfillmentType) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "moveOrder"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(new BasicNameValuePair("orderNumber", orderNumber));
        queryParams.add(new BasicNameValuePair("rate", String.valueOf(rate)));
        amount.ifPresent(amountValue -> queryParams.add(new BasicNameValuePair("amount", amountValue.toString())));
        if ("immediateOrCancel".equalsIgnoreCase(fulfillmentType)) {
            // order can be partially or completely filled, but any portion of the order that cannot be filled immediately will be canceled rather than left on the order book
            queryParams.add(new BasicNameValuePair("immediateOrCancel", "1"));
        } else {
            // order will only be placed if no portion of it fills immediately; this guarantees you will never pay the taker fee on any part of the order that fills
            queryParams.add(new BasicNameValuePair("postOnly", "1"));
        }
        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            post.releaseConnection();
        }
    }

    public Map placeOrder(PoloniexPair pair, BigDecimal rate, BigDecimal amount, String orderType, String fulfillmentType) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(6);
        queryParams.add(new BasicNameValuePair("command", "buy".equalsIgnoreCase(orderType) ? "buy" : "sell"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(new BasicNameValuePair("currencyPair", pair.getPairName()));
        queryParams.add(new BasicNameValuePair("rate", String.valueOf(rate)));
        queryParams.add(new BasicNameValuePair("amount", String.valueOf(amount)));
        if ("fillOrKill".equalsIgnoreCase(fulfillmentType)) {
            // order will either fill in its entirety or be completely aborted
            queryParams.add(new BasicNameValuePair("fillOrKill", "1"));
        } else if ("immediateOrCancel".equalsIgnoreCase(fulfillmentType)) {
            // order can be partially or completely filled, but any portion of the order that cannot be filled immediately will be canceled rather than left on the order book
            queryParams.add(new BasicNameValuePair("immediateOrCancel", "1"));
        } else {
            // order will only be placed if no portion of it fills immediately; this guarantees you will never pay the taker fee on any part of the order that fills
            queryParams.add(new BasicNameValuePair("postOnly", "1"));
        }
        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            post.releaseConnection();
        }
    }

    private HttpPost generatePost(List<NameValuePair> queryParams) throws Exception {
        return HttpUtil.generatePost(getTradingUrl(), queryParams, getApiKey(), getSecretKey(), EncryptionType.HMAC_SHA512);
    }

    private String getTradingUrl() {
        return getBaseUrl() + "tradingApi";
    }
}

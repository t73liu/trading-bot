package io.github.t73liu.exchange.poloniex;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.model.CandlestickIntervals;
import io.github.t73liu.model.currency.PoloniexPair;
import io.github.t73liu.util.DateUtil;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
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
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import javax.validation.constraints.NotNull;
import java.math.BigDecimal;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;
import java.util.Optional;

import static io.github.t73liu.model.currency.PoloniexPair.*;
import static java.nio.charset.StandardCharsets.UTF_8;

@Service
@ConfigurationProperties(prefix = "poloniex")
public class PoloniexService extends ExchangeService {
    private static final Logger LOGGER = LoggerFactory.getLogger(PoloniexService.class);
    // TODO set fees from getFees method
    public static final double TAKER_FEE = 0.0025;
    public static final double MAKER_FEE = 0.0015;

    private final PoloniexTicker ticker;

    @Autowired
    public PoloniexService(PoloniexTicker ticker) {
        this.ticker = ticker;
    }

    // PUBLIC API
    public Map<String, Map<String, String>> getTickers() throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(1);
        queryParams.add(new BasicNameValuePair("command", "returnTicker"));
        HttpGet get = generateGet(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            get.releaseConnection();
        }
    }

    public Map getTickerValue(PoloniexPair pair) throws Exception {
        // TODO sync up getTickers with subscription?
        return ticker.getTickerValue(pair);
    }

    public Map getOrderBook(PoloniexPair pair) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "returnOrderBook"));
        queryParams.add(new BasicNameValuePair("depth", "10"));
        // Set currencyPair to all to see 95 pairs
        queryParams.add(new BasicNameValuePair("currencyPair", pair.getPairName()));
        HttpGet get = generateGet(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            get.releaseConnection();
        }
    }

    public List<Map<String, Double>> getChartData(PoloniexPair pair, LocalDateTime startDateTime, LocalDateTime endDateTime, CandlestickIntervals period) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "returnChartData"));
        // Candlestick period in seconds 300,900,1800,7200,14400,86400
        queryParams.add(new BasicNameValuePair("period", String.valueOf(period.getInterval())));
        queryParams.add(new BasicNameValuePair("currencyPair", pair.getPairName()));
        // UNIX timestamp format of specified time range (i.e. last hour)
        queryParams.add(new BasicNameValuePair("start", String.valueOf(DateUtil.convertToUnixTimestamp(startDateTime))));
        queryParams.add(new BasicNameValuePair("end", String.valueOf(DateUtil.convertToUnixTimestamp(endDateTime))));
        HttpGet get = generateGet(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return mapper.readValue(response.getEntity().getContent(), mapper.getTypeFactory().constructCollectionType(List.class, Map.class));
        } finally {
            get.releaseConnection();
        }
    }

    public Map getCurrencies() throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(1);
        queryParams.add(new BasicNameValuePair("command", "returnCurrencies"));
        HttpGet get = generateGet(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            get.releaseConnection();
        }
    }

    public Map getLoanOrders(String currency) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(2);
        queryParams.add(new BasicNameValuePair("command", "returnLoanOrders"));
        queryParams.add(new BasicNameValuePair("currency", currency));
        HttpGet get = generateGet(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            get.releaseConnection();
        }
    }

    // TODO implement actual arbitrage checker
    public boolean checkArbitrage() throws Exception {
        // Need to check volume of lowest ask and volume of coins in question, low liquidity safer?
        Map<String, Map<String, String>> tickers = getTickers();
        Double btc = Double.valueOf(tickers.get(USDT_BTC.getPairName()).get("lowestAsk"));
        Double eth = Double.valueOf(tickers.get(USDT_ZEC.getPairName()).get("lowestAsk"));
        Double eth2btc = Double.valueOf(tickers.get(BTC_ZEC.getPairName()).get("lowestAsk"));
        double revenue = Math.abs((eth2btc / (eth / btc)) - 1);
        double cost = TAKER_FEE * 3;
        LOGGER.info("exchange: {}, actual: {}, revenue: {}, cost:{}", eth / btc, eth2btc, revenue, cost);
        return revenue > cost;
    }

    // PRIVATE API - NOTE: LIMIT SIX CALLS PER SECOND
    public Map getCompleteBalances() throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(2);
        queryParams.add(new BasicNameValuePair("command", "returnCompleteBalances"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        // Setting account type to all will include margin and lending accounts
//        queryParams.add(new BasicNameValuePair("account", "all"));
        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            post.releaseConnection();
        }
    }

    public Map getDepositAddresses() throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(2);
        queryParams.add(new BasicNameValuePair("command", "returnDepositAddresses"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            post.releaseConnection();
        }
    }

    public Object getOpenOrders(PoloniexPair pair) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "returnOpenOrders"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(pair == null ? new BasicNameValuePair("currencyPair", "all") : new BasicNameValuePair("currencyPair", pair.getPairName()));
        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return mapper.readValue(response.getEntity().getContent(), Object.class);
        } finally {
            post.releaseConnection();
        }
    }

    public Object getTradeHistory(PoloniexPair pair, LocalDateTime startDateTime, LocalDateTime endDateTime) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(5);
        queryParams.add(new BasicNameValuePair("command", "returnTradeHistory"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(pair == null ? new BasicNameValuePair("currencyPair", "all") : new BasicNameValuePair("currencyPair", pair.getPairName()));
        queryParams.add(new BasicNameValuePair("start", String.valueOf(DateUtil.convertToUnixTimestamp(startDateTime))));
        queryParams.add(new BasicNameValuePair("end", String.valueOf(DateUtil.convertToUnixTimestamp(endDateTime))));
        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return mapper.readValue(response.getEntity().getContent(), Object.class);
        } finally {
            post.releaseConnection();
        }
    }

    public Object placeOrder(@NotNull PoloniexPair pair, BigDecimal rate, BigDecimal amount, String orderType, String fulfillmentType) throws Exception {
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
            return mapper.readValue(response.getEntity().getContent(), Object.class);
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
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            post.releaseConnection();
        }
    }

    private Map moveOrder(String orderNumber, double rate, Double amount, String fulfillmentType) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "moveOrder"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(new BasicNameValuePair("orderNumber", orderNumber));
        queryParams.add(new BasicNameValuePair("rate", String.valueOf(rate)));
        if (amount != null) {
            queryParams.add(new BasicNameValuePair("amount", amount.toString()));
        }
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
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            post.releaseConnection();
        }
    }

    private Map withdrawCurrency(String currency, double amount, String depositAddress) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "withdraw"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(new BasicNameValuePair("currency", currency));
        queryParams.add(new BasicNameValuePair("amount", String.valueOf(amount)));
        queryParams.add(new BasicNameValuePair("address", depositAddress));
        // TODO verify purpose of paymentId
//        queryParams.add(new BasicNameValuePair("paymentId", paymentId));

        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            post.releaseConnection();
        }
    }

    private Map getFees() throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "returnFeeInfo"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));

        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return mapper.readValue(response.getEntity().getContent(), Map.class);
        } finally {
            post.releaseConnection();
        }
    }

    // HELPER
    private HttpPost generatePost(List<NameValuePair> queryParams) throws Exception {
        Optional<String> queryParamStr = queryParams.stream()
                .map(Object::toString)
                .reduce((queryOne, queryTwo) -> queryOne + "&" + queryTwo);
        if (!queryParamStr.isPresent()) {
            throw new IllegalArgumentException("Unable to generate query params for Poloniex API: " + queryParams);
        }

        // Generating special headers
        Mac shaMac = Mac.getInstance("HmacSHA512");
        SecretKeySpec keySpec = new SecretKeySpec(getSecretKey().getBytes(UTF_8), "HmacSHA512");
        shaMac.init(keySpec);
        String sign = Hex.encodeHexString(shaMac.doFinal(queryParamStr.get().getBytes(UTF_8)));
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

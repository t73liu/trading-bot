package io.github.t73liu.exchange.poloniex.rest;

import io.github.t73liu.exchange.AccountService;
import io.github.t73liu.exchange.PrivateExchangeService;
import io.github.t73liu.model.EncryptionType;
import io.github.t73liu.model.poloniex.PoloniexPair;
import io.github.t73liu.util.DateUtil;
import io.github.t73liu.util.HttpUtil;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;

import static io.github.t73liu.util.MapperUtil.JSON_READER;

@Service
@ConfigurationProperties(prefix = "poloniex")
public class PoloniexAccountService extends PrivateExchangeService implements AccountService {
    private static final Logger LOGGER = LoggerFactory.getLogger(PoloniexAccountService.class);
    private static final double TAKER_FEE = 0.0025;
    private static final double MAKER_FEE = 0.0015;

    // PRIVATE API - NOTE: LIMIT SIX CALLS PER SECOND
    public Map getCompleteBalances() throws Exception {
        // Setting account to all will include margin and lending accounts
        List<NameValuePair> queryParams = new ObjectArrayList<>(2);
        queryParams.add(new BasicNameValuePair("command", "returnCompleteBalances"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));

        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            post.releaseConnection();
        }
    }

    public Map<String, String> getDepositAddresses() throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(2);
        queryParams.add(new BasicNameValuePair("command", "returnDepositAddresses"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));

        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            post.releaseConnection();
        }
    }

    public Map getTradeHistory(PoloniexPair pair, LocalDateTime startDateTime, LocalDateTime endDateTime) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(5);
        queryParams.add(new BasicNameValuePair("command", "returnTradeHistory"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(pair == null ? new BasicNameValuePair("currencyPair", "all") : new BasicNameValuePair("currencyPair", pair.getPairName()));
        queryParams.add(new BasicNameValuePair("start", String.valueOf(DateUtil.localDateTimeToUnixTimestamp(startDateTime))));
        queryParams.add(new BasicNameValuePair("end", String.valueOf(DateUtil.localDateTimeToUnixTimestamp(endDateTime))));

        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            post.releaseConnection();
        }
    }

    private Map withdrawCurrency(String currency, double amount, String depositAddress) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(5);
        // TODO investigate paymentId parameter
        queryParams.add(new BasicNameValuePair("command", "withdraw"));
        queryParams.add(new BasicNameValuePair("nonce", String.valueOf(System.currentTimeMillis())));
        queryParams.add(new BasicNameValuePair("currency", currency));
        queryParams.add(new BasicNameValuePair("amount", String.valueOf(amount)));
        queryParams.add(new BasicNameValuePair("address", depositAddress));

        HttpPost post = generatePost(queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(post)) {
            return JSON_READER.readValue(response.getEntity().getContent());
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
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            post.releaseConnection();
        }
    }

    private HttpPost generatePost(List<NameValuePair> queryParams) throws Exception {
        return HttpUtil.generatePost(getAccountUrl(), queryParams, getApiKey(), getSecretKey(), EncryptionType.HMAC_SHA512);
    }

    private String getAccountUrl() {
        return getBaseUrl() + "tradingApi";
    }
}

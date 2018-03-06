package io.github.t73liu.news;

import io.github.t73liu.exchange.PrivateExchangeService;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import java.util.List;

import static io.github.t73liu.util.HttpUtil.generateGet;
import static io.github.t73liu.util.MapperUtil.JSON_READER;

@Service
@ConfigurationProperties(prefix = "alphavantage")
public class AlphaVantageService extends PrivateExchangeService {
    public Object getCandle(String symbol) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(5);
        queryParams.add(new BasicNameValuePair("function", "TIME_SERIES_INTRADAY"));
        queryParams.add(new BasicNameValuePair("symbol", symbol));
        // following values are supported: 1min, 5min, 15min, 30min, 60min
        queryParams.add(new BasicNameValuePair("interval", "1min"));
        // compact returns only the latest 100 data points in the intraday time series
        // full returns the full-length intraday time series
        queryParams.add(new BasicNameValuePair("outputsize", "compact"));
        queryParams.add(new BasicNameValuePair("apikey", getApiKey()));
        HttpGet get = generateGet(getBaseUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    public Object getEMA(String symbol, int numOfEntries) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(6);
        queryParams.add(new BasicNameValuePair("function", "EMA"));
        queryParams.add(new BasicNameValuePair("symbol", symbol));
        // following values are supported: 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly
        queryParams.add(new BasicNameValuePair("interval", "1min"));
        queryParams.add(new BasicNameValuePair("time_period", String.valueOf(numOfEntries)));
        // Four types are supported: close, open, high, low
        queryParams.add(new BasicNameValuePair("series_type", "close"));
        queryParams.add(new BasicNameValuePair("apikey", getApiKey()));
        HttpGet get = generateGet(getBaseUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    public Object getMACD(String symbol, int fastPeriod, int slowPeriod, int signalPeriod) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(6);
        queryParams.add(new BasicNameValuePair("function", "EMA"));
        queryParams.add(new BasicNameValuePair("symbol", symbol));
        // following values are supported: 1min, 5min, 15min, 30min, 60min, daily, weekly, monthly
        queryParams.add(new BasicNameValuePair("interval", "1min"));
        // Four types are supported: close, open, high, low
        queryParams.add(new BasicNameValuePair("series_type", "close"));
        queryParams.add(new BasicNameValuePair("fastperiod", String.valueOf(fastPeriod)));
        queryParams.add(new BasicNameValuePair("slowperiod", String.valueOf(slowPeriod)));
        queryParams.add(new BasicNameValuePair("signalperiod", String.valueOf(signalPeriod)));
        queryParams.add(new BasicNameValuePair("apikey", getApiKey()));
        HttpGet get = generateGet(getBaseUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }
}

package io.github.t73liu.exchange.poloniex.rest;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.exchange.MarketService;
import io.github.t73liu.model.poloniex.*;
import it.unimi.dsi.fastutil.objects.Object2ObjectMaps.UnmodifiableMap;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;
import org.ta4j.core.Bar;

import java.util.Arrays;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

import static io.github.t73liu.util.HttpUtil.generateGet;
import static io.github.t73liu.util.MapperUtil.JSON_READER;
import static io.github.t73liu.util.MapperUtil.TYPE_FACTORY;

@Service
@ConfigurationProperties(prefix = "poloniex")
public class PoloniexMarketService extends ExchangeService implements MarketService {
    private static List<NameValuePair> getCandleParameters(PoloniexPair pair, long startSeconds, long endSeconds, PoloniexCandleInterval period) {
        List<NameValuePair> queryParams = new ObjectArrayList<>(5);
        queryParams.add(new BasicNameValuePair("command", "returnChartData"));
        queryParams.add(new BasicNameValuePair("period", period.getIntervalName()));
        queryParams.add(new BasicNameValuePair("currencyPair", pair.getPairName()));
        queryParams.add(new BasicNameValuePair("start", String.valueOf(startSeconds)));
        queryParams.add(new BasicNameValuePair("end", String.valueOf(endSeconds)));
        return queryParams;
    }

    public PoloniexTicker getTickerForPair(PoloniexPair pair) throws Exception {
        return getAllTicker().get(pair);
    }

    public Map<PoloniexPair, PoloniexTicker> getAllTicker() throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(1);
        queryParams.add(new BasicNameValuePair("command", "returnTicker"));
        HttpGet get = generateGet(getPublicUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(TYPE_FACTORY.constructMapType(UnmodifiableMap.class, PoloniexPair.class, PoloniexTicker.class))
                    .readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    public Map<PoloniexPair, PoloniexOrderBook> getAllOrderBook(int amount) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "returnOrderBook"));
        queryParams.add(new BasicNameValuePair("depth", String.valueOf(amount)));
        queryParams.add(new BasicNameValuePair("currencyPair", "all"));
        HttpGet get = generateGet(getPublicUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(TYPE_FACTORY.constructMapType(UnmodifiableMap.class, PoloniexPair.class, PoloniexOrderBook.class))
                    .readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    public PoloniexOrderBook getOrderBookForPair(PoloniexPair pair, int amount) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(3);
        queryParams.add(new BasicNameValuePair("command", "returnOrderBook"));
        queryParams.add(new BasicNameValuePair("depth", String.valueOf(amount)));
        queryParams.add(new BasicNameValuePair("currencyPair", pair.getPairName()));
        HttpGet get = generateGet(getPublicUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(PoloniexOrderBook.class).readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    public List<Bar> getCandlestickForPair(PoloniexPair pair, long startSeconds, long endSeconds, PoloniexCandleInterval period) throws Exception {
        return Arrays.stream(getExchangeCandleForPair(pair, startSeconds, endSeconds, period))
                .map(PoloniexCandle::toTick)
                .collect(Collectors.toCollection(ObjectArrayList::new));
    }

    private PoloniexCandle[] getExchangeCandleForPair(PoloniexPair pair, long startSeconds, long endSeconds, PoloniexCandleInterval period) throws Exception {
        HttpGet get = generateGet(getPublicUrl(), getCandleParameters(pair, startSeconds, endSeconds, period));

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(PoloniexCandle[].class).readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    private String getPublicUrl() {
        return getBaseUrl() + "public";
    }
}

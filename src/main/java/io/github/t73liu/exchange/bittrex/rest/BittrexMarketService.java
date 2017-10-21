package io.github.t73liu.exchange.bittrex.rest;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.exchange.MarketService;
import io.github.t73liu.model.bittrex.BittrexOrderBook;
import io.github.t73liu.model.bittrex.BittrexPair;
import io.github.t73liu.model.bittrex.BittrexTicker;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import java.util.List;

import static io.github.t73liu.util.HttpUtil.generateGet;
import static io.github.t73liu.util.MapperUtil.JSON_READER;

@Service
@ConfigurationProperties(prefix = "bittrex")
public class BittrexMarketService extends ExchangeService implements MarketService {
    private static final Logger LOGGER = LoggerFactory.getLogger(BittrexMarketService.class);
    private static final double FEE = 0.0025;

    public BittrexTicker getTickerForPair(BittrexPair pair) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(1);
        queryParams.add(new BasicNameValuePair("market", pair.getPairName()));
        HttpGet get = generateGet(getTickerUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(BittrexTicker.class).readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    public BittrexOrderBook getOrderBookForPair(BittrexPair pair) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(2);
        queryParams.add(new BasicNameValuePair("market", pair.getPairName()));
        // type can be buy, sell or both
        queryParams.add(new BasicNameValuePair("type", "both"));
        HttpGet get = generateGet(getOrderBookUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(BittrexOrderBook.class).readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    private String getTickerUrl() {
        return getBaseUrl() + "public/getticker";
    }

    private String getOrderBookUrl() {
        return getBaseUrl() + "public/getorderbook";
    }
}

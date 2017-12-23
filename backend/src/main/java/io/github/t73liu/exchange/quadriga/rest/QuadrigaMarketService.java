package io.github.t73liu.exchange.quadriga.rest;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.exchange.MarketService;
import io.github.t73liu.model.quadriga.QuadrigaOrderBook;
import io.github.t73liu.model.quadriga.QuadrigaPair;
import io.github.t73liu.model.quadriga.QuadrigaTicker;
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
@ConfigurationProperties(prefix = "quadrigacx")
public class QuadrigaMarketService extends ExchangeService implements MarketService {
    private static final Logger LOGGER = LoggerFactory.getLogger(QuadrigaMarketService.class);
    private static final double FIAT_FEE = 0.005;
    private static final double CRYPTO_FEE = 0.002;

    public QuadrigaTicker getTickerForPair(QuadrigaPair pair) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(1);
        queryParams.add(new BasicNameValuePair("book", pair.getPairName()));
        HttpGet get = generateGet(getTickerUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(QuadrigaTicker.class).readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    private String getTickerUrl() {
        return getBaseUrl() + "ticker";
    }


    public QuadrigaOrderBook getOrderBookForPair(QuadrigaPair pair) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(1);
        queryParams.add(new BasicNameValuePair("book", pair.getPairName()));
        HttpGet get = generateGet(getOrderBookUrl(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(QuadrigaOrderBook.class).readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    private String getOrderBookUrl() {
        return getBaseUrl() + "order_book";
    }
}

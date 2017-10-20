package io.github.t73liu.exchange.bitfinex.rest;

import io.github.t73liu.exchange.ExchangeService;
import io.github.t73liu.exchange.MarketService;
import io.github.t73liu.model.bitfinex.BitfinexPair;
import io.github.t73liu.model.bitfinex.BitfinexTicker;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import java.util.List;

import static io.github.t73liu.util.HttpUtil.generateGet;
import static io.github.t73liu.util.MapperUtil.JSON_READER;

@Service
@ConfigurationProperties(prefix = "bitfinex")
public class BitfinexMarketService extends ExchangeService implements MarketService {
    private String TICKER_URL;

    @PostConstruct
    public void init() {
        TICKER_URL = getBaseUrl() + "ticker/";
    }

    public BitfinexTicker getTickerForPair(BitfinexPair pair) throws Exception {
        return new BitfinexTicker(getExchangeTickerForPair(pair));
    }

    private double[] getExchangeTickerForPair(BitfinexPair pair) throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(0);
        HttpGet get = generateGet(TICKER_URL + pair.getPairName(), queryParams);

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return JSON_READER.forType(double[].class).readValue(response.getEntity().getContent());
        } finally {
            get.releaseConnection();
        }
    }

    // TODO candles and trades
    // https://api.bitfinex.com/v2/candles/trade:5m:tBTCUSD/last
    // https://api.bitfinex.com/v2/candles/trade:1m:tBTCUSD/hist
    // https://api.bitfinex.com/v2/trades/tBTCUSD
}

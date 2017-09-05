package io.github.t73liu.service;

import com.fasterxml.jackson.databind.node.ArrayNode;
import io.github.t73liu.model.currency.PoloniexPair;
import io.github.t73liu.util.WampClientFactory;
import it.unimi.dsi.fastutil.objects.Object2DoubleOpenHashMap;
import it.unimi.dsi.fastutil.objects.Object2ObjectOpenHashMap;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;
import ws.wamp.jawampa.PubSubData;
import ws.wamp.jawampa.WampClient;
import ws.wamp.jawampa.WampClient.ConnectedState;

import javax.annotation.PostConstruct;
import java.util.Map;

import static io.github.t73liu.model.currency.PoloniexPair.ETH_ETC;

@Service
public class PoloniexTicker {
    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());

    private WampClient client;

    private Map<PoloniexPair, Map<String, Double>> tickerMap = new Object2ObjectOpenHashMap<>(PoloniexPair.values().length);

    private PoloniexTicker() throws Exception {
        this.client = WampClientFactory.getNewInstance("wss://api.poloniex.com", "realm1");
    }

    public WampClient getClient() {
        return client;
    }

    public Map<String, Double> getTickerValue(PoloniexPair pair) {
        return tickerMap.get(pair);
    }

    @PostConstruct
    private void init() {
        LOGGER.info("Opening WAMP Client and subscribing to ticker");
        this.client.statusChanged().subscribe(state -> {
            if (state instanceof ConnectedState) {
                client.makeSubscription("ticker")
                        .subscribe(this::updateTicker);
//                client.makeSubscription("ETH_ETC")
//                        .subscribe(s -> System.out.println(s.arguments()));
            }
        });
        this.client.open();
    }

    private void updateTicker(PubSubData subData) {
        // currencyPair, last, lowestAsk, highestBid, percentChange, baseVolume, quoteVolume, isFrozen, 24hrHigh, 24hrLow
        ArrayNode array = subData.arguments();
        String currency = array.get(0).asText();
        LOGGER.info("Updating {} with ticker: {}", currency, array);
        if (ETH_ETC.getPairName().equals(currency)) {
            PoloniexPair pair = PoloniexPair.valueOf(currency);
            // TODO hardcoded ticker size and map, need to make a ticker pojo
            Map<String, Double> ticker = new Object2DoubleOpenHashMap<>(9);
            ticker.put("last", array.get(1).asDouble());
            ticker.put("lowestAsk", array.get(2).asDouble());
            ticker.put("highestBid", array.get(3).asDouble());
            ticker.put("percentChange", array.get(4).asDouble());
            ticker.put("baseVolume", array.get(5).asDouble());
            ticker.put("quoteVolume", array.get(6).asDouble());
            ticker.put("isFrozen", array.get(7).asDouble());
            ticker.put("24hrHigh", array.get(8).asDouble());
            ticker.put("24hrLow", array.get(9).asDouble());
            tickerMap.put(pair, ticker);
        }
    }
}

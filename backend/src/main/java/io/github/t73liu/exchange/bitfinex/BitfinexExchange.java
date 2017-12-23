package io.github.t73liu.exchange.bitfinex;

import io.github.t73liu.exchange.bitfinex.websocket.BitfinexSocket;
import org.eclipse.jetty.websocket.client.WebSocketClient;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import java.net.URI;

@Service
public class BitfinexExchange {
    @PostConstruct
    private void initializeSocket() throws Exception {
        WebSocketClient client = new WebSocketClient();
        client.start();
        client.connect(new BitfinexSocket(), new URI("wss://api.bitfinex.com/ws/2"));
    }
}

package io.github.t73liu.exchange.bitfinex.websocket;

import com.google.common.collect.ImmutableMap;
import org.eclipse.jetty.websocket.api.Session;
import org.eclipse.jetty.websocket.api.annotations.OnWebSocketClose;
import org.eclipse.jetty.websocket.api.annotations.OnWebSocketConnect;
import org.eclipse.jetty.websocket.api.annotations.OnWebSocketMessage;
import org.eclipse.jetty.websocket.api.annotations.WebSocket;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.util.concurrent.CountDownLatch;

import static io.github.t73liu.util.MapperUtil.JSON_WRITER;

@WebSocket(maxTextMessageSize = 64 * 1024)
public class BitfinexSocket {
    private static final Logger LOGGER = LoggerFactory.getLogger(BitfinexSocket.class);
    private final CountDownLatch closeLatch = new CountDownLatch(1);
    private Session session;

    @OnWebSocketConnect
    public void onConnect(Session session) throws IOException {
        LOGGER.info("Connecting to socket session: {}", session);
        this.session = session;
        if (session != null) {
            session.getRemote().sendString(JSON_WRITER.writeValueAsString(ImmutableMap.of("event", "subscribe", "channel", "ticker", "symbol", "tBTCUSD")));
            session.getRemote().sendString(JSON_WRITER.writeValueAsString(ImmutableMap.of("event", "subscribe", "channel", "candles", "key", "trade:1m:tBTCUSD")));
        }
    }

    @OnWebSocketMessage
    public void onMessage(String message) {
        LOGGER.info("Message Received: {}", message);
    }

    @OnWebSocketClose
    public void onClose(int statusCode, String reason) {
        LOGGER.info("Closing Session with code: {}, reason: {}", statusCode, reason);
        this.session.close();
        this.closeLatch.countDown();
    }
}

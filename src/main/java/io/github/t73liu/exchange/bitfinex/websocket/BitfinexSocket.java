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
import java.io.InputStream;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

import static io.github.t73liu.util.MapperUtil.JSON_WRITER;

@WebSocket(maxTextMessageSize = 64 * 1024)
public class BitfinexSocket {
    private static final Logger LOGGER = LoggerFactory.getLogger(BitfinexSocket.class);
    private final CountDownLatch closeLatch;
    private Session session;

    public BitfinexSocket() {
        this.closeLatch = new CountDownLatch(1);
    }

    public boolean awaitClose(int duration, TimeUnit unit) throws InterruptedException {
        return this.closeLatch.await(duration, unit);
    }

    @OnWebSocketClose
    public void onClose(int statusCode, String reason) {
        LOGGER.info("Closing Session with code: {}, reason: {}", statusCode, reason);
        this.session = null;
        this.closeLatch.countDown(); // trigger latch
    }

    @OnWebSocketConnect
    public void onConnect(Session session) throws IOException {
        LOGGER.info("Connecting to socket session: {}", session);
        this.session = session;
        if (session != null) {
            session.getRemote().sendStringByFuture(JSON_WRITER.writeValueAsString(ImmutableMap.of("event", "subscribe", "channel", "ticker", "symbol", "tBTCUSD")));
            session.getRemote().sendStringByFuture(JSON_WRITER.writeValueAsString(ImmutableMap.of("event", "subscribe", "channel", "candles", "key", "trade:1m:tBTCUSD")));
        }
    }

    @OnWebSocketMessage
    public void onMessage(InputStream inputStream) throws Exception {
        LOGGER.info("Message Received: {}", JSON_WRITER.writeValueAsString(inputStream));
    }
}

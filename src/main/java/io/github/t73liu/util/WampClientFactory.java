package io.github.t73liu.util;

import ws.wamp.jawampa.WampClient;
import ws.wamp.jawampa.WampClientBuilder;
import ws.wamp.jawampa.connection.IWampConnectorProvider;
import ws.wamp.jawampa.transport.netty.NettyWampClientConnectorProvider;

import java.util.concurrent.TimeUnit;

public class WampClientFactory {
    public static WampClient getNewInstance(String uri, String realm) throws Exception {
        WampClientBuilder builder = new WampClientBuilder();
        IWampConnectorProvider connectorProvider = new NettyWampClientConnectorProvider();
        builder.withConnectorProvider(connectorProvider)
                .withUri(uri)
                .withRealm(realm)
                .withInfiniteReconnects()
                .withReconnectInterval(10, TimeUnit.SECONDS);
        return builder.build();
    }
}

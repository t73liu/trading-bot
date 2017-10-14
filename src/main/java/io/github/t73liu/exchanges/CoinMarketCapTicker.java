package io.github.t73liu.exchanges;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.github.t73liu.util.ObjectMapperFactory;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.utils.URIBuilder;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import java.util.List;
import java.util.Map;

@Service
public class CoinMarketCapTicker {
    private ObjectMapper mapper;

    @PostConstruct
    private void init() {
        this.mapper = ObjectMapperFactory.getNewInstance();
    }

    public List<Map<String, String>> getTickers() throws Exception {
        List<NameValuePair> queryParams = new ObjectArrayList<>(1);
        queryParams.add(new BasicNameValuePair("limit", "200"));

        HttpGet get = new HttpGet(new URIBuilder("https://api.coinmarketcap.com/v1/ticker/").addParameters(queryParams).build());

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return mapper.readValue(response.getEntity().getContent(), mapper.getTypeFactory().constructCollectionType(List.class, Map.class));
        } finally {
            get.releaseConnection();
        }
    }
}

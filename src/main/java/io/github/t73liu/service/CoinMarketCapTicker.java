package io.github.t73liu.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.github.t73liu.util.ObjectMapperFactory;
import org.apache.http.NameValuePair;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.utils.URIBuilder;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClients;
import org.apache.http.message.BasicNameValuePair;
import org.springframework.stereotype.Service;

import javax.annotation.PostConstruct;
import java.util.ArrayList;
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
        List<NameValuePair> queryParams = new ArrayList<>();
        queryParams.add(new BasicNameValuePair("limit", "10"));

        HttpGet get = new HttpGet(new URIBuilder("https://api.coinmarketcap.com/v1/ticker/").addParameters(queryParams).build());

        try (CloseableHttpClient httpClient = HttpClients.createDefault();
             CloseableHttpResponse response = httpClient.execute(get)) {
            return mapper.readValue(response.getEntity().getContent(), mapper.getTypeFactory().constructCollectionType(List.class, Map.class));
        } finally {
            get.releaseConnection();
        }
    }
}

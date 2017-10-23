package io.github.t73liu.exchange.bitfinex.rest;

import io.github.t73liu.exchange.OrderService;
import io.github.t73liu.exchange.PrivateExchangeService;
import io.github.t73liu.model.EncryptionType;
import org.apache.commons.codec.binary.Hex;
import org.apache.http.NameValuePair;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.HttpPost;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.util.List;
import java.util.stream.Collectors;

import static io.github.t73liu.util.MapperUtil.JSON_WRITER;
import static java.nio.charset.StandardCharsets.UTF_8;

@Service
@ConfigurationProperties(prefix = "bitfinex")
public class BitfinexOrderService extends PrivateExchangeService implements OrderService {
    // TODO implement
    public HttpPost generatePost(String apiPath, List<NameValuePair> queryParams, String apiKey, String secretKey) throws Exception {
        String queryParamStr = JSON_WRITER.writeValueAsString(queryParams.stream()
                .collect(Collectors.toMap(NameValuePair::getName, NameValuePair::getValue)));
        String signature = "/api/" + apiPath + System.currentTimeMillis() + queryParamStr;
        EncryptionType encryptionType = EncryptionType.HMAC_SHA384;

        Mac shaMac = Mac.getInstance(encryptionType.getName());
        shaMac.init(new SecretKeySpec(secretKey.getBytes(UTF_8), encryptionType.getName()));
        String encryptedSignature = Hex.encodeHexString(shaMac.doFinal(signature.getBytes(UTF_8)));
        HttpPost post = new HttpPost(getBaseUrl() + apiPath);
        post.addHeader("bfx-nonce", String.valueOf(System.currentTimeMillis()));
        post.addHeader("bfx-apikey", apiKey);
        post.addHeader("bfx-signature", encryptedSignature);
        post.setEntity(new UrlEncodedFormEntity(queryParams));
        return post;
    }
}

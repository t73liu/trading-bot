package io.github.t73liu.util;

import io.github.t73liu.model.EncryptionType;
import org.apache.commons.codec.binary.Hex;
import org.apache.http.NameValuePair;
import org.apache.http.client.entity.UrlEncodedFormEntity;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.client.utils.URIBuilder;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.util.List;

import static java.nio.charset.StandardCharsets.UTF_8;

public class HttpUtil {
    public static HttpGet generateGet(String url, List<NameValuePair> queryParams) throws Exception {
        return new HttpGet(new URIBuilder(url).addParameters(queryParams).build());
    }

    public static HttpPost generatePost(String url, List<NameValuePair> queryParams, String apiKey, String secretKey, EncryptionType encryptionType) throws Exception {
        String queryParamStr = queryParams.stream()
                .map(Object::toString)
                .reduce((queryOne, queryTwo) -> queryOne + "&" + queryTwo)
                .orElse("");

        // Generating special headers
        Mac shaMac = Mac.getInstance(encryptionType.getName());
        shaMac.init(new SecretKeySpec(secretKey.getBytes(UTF_8), encryptionType.getName()));
        String sign = Hex.encodeHexString(shaMac.doFinal(queryParamStr.getBytes(UTF_8)));
        HttpPost post = new HttpPost(url);
        post.addHeader("Key", apiKey);
        post.addHeader("Sign", sign);
        post.setEntity(new UrlEncodedFormEntity(queryParams));
        return post;
    }
}

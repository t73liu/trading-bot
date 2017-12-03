package io.github.t73liu.news;

import com.rometools.rome.feed.synd.SyndEntry;
import com.rometools.rome.feed.synd.SyndFeed;
import com.rometools.rome.io.SyndFeedInput;
import com.rometools.rome.io.XmlReader;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.springframework.stereotype.Component;

import java.net.URL;
import java.util.List;

@Component
public class RssService {
    private List<String> subscriptions = new ObjectArrayList<>();

    public List<SyndEntry> getEntries(String url) throws Exception {
        try (XmlReader reader = new XmlReader(new URL(url))) {
            SyndFeed feed = new SyndFeedInput().build(reader);
            return feed.getEntries();
        }
    }

    public Object getLatest() throws Exception {
        return getEntries("https://www.reddit.com/r/CryptoMarkets+CryptoCurrency/top/.rss?sort=new");
    }

    public List<String> getSubscriptions() {
        return subscriptions;
    }
}

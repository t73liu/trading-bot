package io.github.t73liu.news;

import com.rometools.rome.feed.synd.SyndEntry;
import com.rometools.rome.feed.synd.SyndFeed;
import com.rometools.rome.io.SyndFeedInput;
import com.rometools.rome.io.XmlReader;
import it.unimi.dsi.fastutil.objects.Object2ObjectOpenHashMap;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

import java.net.URL;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@Component
@ConfigurationProperties("news.rss")
public class RssService {
    private List<String> subscriptions = new ObjectArrayList<>();

    public List<String> getSubscriptions() {
        return subscriptions;
    }

    public Object getLatest() throws Exception {
        Map<String, List<String>> latestNews = new Object2ObjectOpenHashMap<>(subscriptions.size());
        for (String feed : subscriptions) {
            latestNews.put(feed, getEntries(feed).stream().map(SyndEntry::getTitle).collect(Collectors.toList()));
        }
        return latestNews;
    }

    private static List<SyndEntry> getEntries(String url) throws Exception {
        try (XmlReader reader = new XmlReader(new URL(url))) {
            SyndFeed feed = new SyndFeedInput().build(reader);
            return feed.getEntries();
        }
    }
}

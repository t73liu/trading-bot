package io.github.t73liu.news;

import com.google.common.primitives.Longs;
import it.unimi.dsi.fastutil.longs.LongOpenHashSet;
import it.unimi.dsi.fastutil.objects.ObjectArraySet;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;
import twitter4j.*;
import twitter4j.api.TweetsResources;
import twitter4j.api.UsersResources;
import twitter4j.conf.Configuration;
import twitter4j.conf.ConfigurationBuilder;

import javax.annotation.PostConstruct;
import java.util.Set;

@Component
@ConfigurationProperties("news.twitter")
public class TwitterService {
    private static final Logger LOGGER = LoggerFactory.getLogger(TwitterService.class);

    @Value("${news.twitter.consumerKey}")
    private String consumerKey;

    @Value("${news.twitter.consumerSecret}")
    private String consumerSecret;

    @Value("${news.twitter.accessToken}")
    private String accessToken;

    @Value("${news.twitter.accessTokenSecret}")
    private String accessTokenSecret;

    private Set<String> users = new ObjectArraySet<>();

    public Set<String> getUsers() {
        return users;
    }

    public Object getLatest() {
        return null;
    }

    @PostConstruct
    private void initializeStream() {
        Configuration configuration = new ConfigurationBuilder()
                .setDebugEnabled(true)
                .setOAuthConsumerKey(consumerKey)
                .setOAuthConsumerSecret(consumerSecret)
                .setOAuthAccessToken(accessToken)
                .setOAuthAccessTokenSecret(accessTokenSecret)
                .build();
        Twitter twitter = new TwitterFactory(configuration).getInstance();
        UsersResources usersResources = twitter.users();
        TweetsResources tweetsResources = twitter.tweets();
        Set<Long> userIds = new LongOpenHashSet();
        for (String userName : users) {
            try {
                userIds.add(usersResources.showUser(userName).getId());
            } catch (TwitterException e) {
                LOGGER.warn("Unable to lookup user name: {}", userName);
            }
        }
        TwitterStream twitterStream = new TwitterStreamFactory(configuration).getInstance();
        StatusListener listener = new StatusListener() {
            @Override
            public void onStatus(Status status) {
                if (userIds.contains(status.getUser().getId())) {
                    LOGGER.info("User: {}, Text: {}", status.getUser().getScreenName(), status.getText());
                    long inReplyToStatusId = status.getInReplyToStatusId();
                    if (inReplyToStatusId > 0) {
                        try {
                            Status repliedToTweet = tweetsResources.showStatus(inReplyToStatusId);
                            LOGGER.info("User: {}, Text: {}", repliedToTweet.getUser().getScreenName(), repliedToTweet.getText());
                        } catch (TwitterException e) {
                            LOGGER.warn("Unable to lookup repliedToTweet status: {}", inReplyToStatusId);
                        }
                    }
                    // TODO verify if required
                    if (status.isRetweet()) {
                        Status retweetedStatus = status.getRetweetedStatus();
                        LOGGER.info("Retweeted By: {} - User: {}, Text: {}", status.getUser().getScreenName(), retweetedStatus.getUser().getScreenName(), retweetedStatus.getText());
                    }
                }
            }

            @Override
            public void onDeletionNotice(StatusDeletionNotice statusDeletionNotice) {
                LOGGER.info("Got a status deletion notice id: {}", statusDeletionNotice.getStatusId());
            }

            @Override
            public void onTrackLimitationNotice(int numberOfLimitedStatuses) {
                LOGGER.info("Got track limitation notice: {}", numberOfLimitedStatuses);
            }

            @Override
            public void onScrubGeo(long userId, long upToStatusId) {
                LOGGER.info("Got scrub_geo event userId: {}, upToStatusId: {}", userId, upToStatusId);
            }

            @Override
            public void onStallWarning(StallWarning warning) {
                LOGGER.warn("Got stall warning: {}", warning);
            }

            @Override
            public void onException(Exception ex) {
                LOGGER.error("Twitter Exception: {}", ex.getMessage(), ex);
            }
        };
        twitterStream.addListener(listener);
        twitterStream.filter(new FilterQuery(Longs.toArray(userIds)).language("en"));
    }
}

package io.github.t73liu.rest;

import com.google.common.collect.ImmutableMap;
import io.github.t73liu.news.RssService;
import io.github.t73liu.news.TwitterService;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.ws.rs.Consumes;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.function.Function;

@Component
@Path("/news")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Tag(name = "NewsResource")
public class NewsResource {
    @Autowired
    private RssService rssSource;
    @Autowired(required = false)
    private TwitterService twitterService;

    @GET
    @Path("/rss")
    @ApiResponse(responseCode = "200", description = "Get Latest Unread News From RSS Feed", content = @Content(schema = @Schema(implementation = Object.class)))
    public Response getRssFeeds() throws Exception {
        return Response.ok(rssSource.getLatest()).build();
    }

    @GET
    @Path("/twitter")
    @ApiResponse(responseCode = "200", description = "Get Latest Unread News From Twitter Feed", content = @Content(schema = @Schema(implementation = Object.class)))
    public Response getTwitterFeeds() {
        return Response.ok(callTwitter(TwitterService::getLatest)).build();
    }

    private Object callTwitter(Function<TwitterService, Object> function) {
        return twitterService != null ? function.apply(twitterService) : ImmutableMap.of("service", "Twitter", "status", "disabled");
    }
}

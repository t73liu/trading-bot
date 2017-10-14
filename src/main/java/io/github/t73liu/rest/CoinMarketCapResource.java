package io.github.t73liu.rest;

import io.github.t73liu.exchanges.CoinMarketCapTicker;
import io.github.t73liu.model.ExceptionWrapper;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiResponse;
import io.swagger.annotations.ApiResponses;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.ws.rs.Consumes;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Map;

@Component
@Path("/coinmarketcap")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("CoinMarketCapResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class CoinMarketCapResource {
    private final CoinMarketCapTicker ticker;

    @Autowired
    public CoinMarketCapResource(CoinMarketCapTicker ticker) {
        this.ticker = ticker;
    }

    @GET
    @Path("/tickers")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Currency Information", responseContainer = "List", response = Map.class))
    public Response test() throws Exception {
        return Response.ok(ticker.getTickers()).build();
    }
}

package io.github.t73liu.rest;

import io.github.t73liu.exchange.bitfinex.rest.BitfinexMarketService;
import io.github.t73liu.model.bitfinex.*;
import io.swagger.v3.oas.annotations.Parameter;
import io.swagger.v3.oas.annotations.media.ArraySchema;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;

@Component
@Path("/bitfinex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Tag(name = "BitfinexResource")
public class BitfinexResource {
    private final BitfinexMarketService marketService;

    @Autowired
    public BitfinexResource(BitfinexMarketService marketService) {
        this.marketService = marketService;
    }

    @GET
    @Path("/tickers/{pair}")
    @ApiResponse(responseCode = "200", description = "Retrieved Ticker of Specified Pair in Bitfinex", content = @Content(schema = @Schema(implementation = BitfinexTicker.class)))
    public Response getTickerForPair(@PathParam("pair") @Valid @NotNull BitfinexPair pair) throws Exception {
        return Response.ok(marketService.getTickerForPair(pair)).build();
    }

    @GET
    @Path("/candles/{pair}")
    @ApiResponse(responseCode = "200", description = "Retrieved Ticker of Specified Pair in Bitfinex", content = @Content(schema = @Schema(implementation = BitfinexCandle.class)))
    public Response getCandleForPair(@Parameter(schema = @Schema(implementation = BitfinexPair.class)) @PathParam("pair") @Valid @NotNull BitfinexPair pair,
                                     @QueryParam("interval") @Valid @NotNull BitfinexCandleInterval candleInterval,
                                     @QueryParam("startMilliseconds") Long startMilliseconds,
                                     @QueryParam("endMilliseconds") Long endMilliseconds,
                                     @QueryParam("limit") @DefaultValue("10") int limit,
                                     @QueryParam("newestFirst") @DefaultValue("true") boolean newestFirst) throws Exception {
        return Response.ok(marketService.getExchangeCandleForPair(pair, candleInterval, startMilliseconds, endMilliseconds, limit, newestFirst)).build();
    }

    @GET
    @Path("/orderBooks/{pair}")
    @ApiResponse(responseCode = "200", description = "Retrieved Order Book of Specified Pair in Bitfinex", content = @Content(array = @ArraySchema(schema = @Schema(implementation = BitfinexOrderBook.class))))
    public Response getOrderBookForPair(@Parameter(schema = @Schema(implementation = BitfinexPair.class)) @PathParam("pair") @Valid @NotNull BitfinexPair pair) throws Exception {
        return Response.ok(marketService.getExchangeOrderBookForPair(pair)).build();
    }
}

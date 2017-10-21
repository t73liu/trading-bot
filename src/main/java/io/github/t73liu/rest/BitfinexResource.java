package io.github.t73liu.rest;

import io.github.t73liu.exception.ExceptionWrapper;
import io.github.t73liu.exchange.bitfinex.rest.BitfinexAccountService;
import io.github.t73liu.exchange.bitfinex.rest.BitfinexMarketService;
import io.github.t73liu.exchange.bitfinex.rest.BitfinexOrderService;
import io.github.t73liu.model.bitfinex.*;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiResponse;
import io.swagger.annotations.ApiResponses;
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
@Api("BitfinexResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class BitfinexResource {
    private final BitfinexAccountService accountService;
    private final BitfinexMarketService marketService;
    private final BitfinexOrderService orderService;

    @Autowired
    public BitfinexResource(BitfinexAccountService accountService, BitfinexMarketService marketService, BitfinexOrderService orderService) {
        this.accountService = accountService;
        this.marketService = marketService;
        this.orderService = orderService;
    }

    @GET
    @Path("/tickers/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Ticker of Specified Pair in Bitfinex", response = BitfinexTicker.class))
    public Response getTickerForPair(@PathParam("pair") @Valid @NotNull BitfinexPair pair) throws Exception {
        return Response.ok(marketService.getTickerForPair(pair)).build();
    }

    @GET
    @Path("/candles/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Ticker of Specified Pair in Bitfinex", response = BitfinexCandle.class))
    public Response getCandleForPair(@PathParam("pair") @Valid @NotNull BitfinexPair pair,
                                     @QueryParam("interval") @Valid @NotNull BitfinexCandleInterval candleInterval,
                                     @QueryParam("startMilliseconds") Long startMilliseconds,
                                     @QueryParam("endMilliseconds") Long endMilliseconds,
                                     @QueryParam("limit") @DefaultValue("10") int limit,
                                     @QueryParam("newestFirst") @DefaultValue("true") boolean newestFirst) throws Exception {
        return Response.ok(marketService.getExchangeCandleForPair(pair, candleInterval, startMilliseconds, endMilliseconds, limit, newestFirst)).build();
    }

    @GET
    @Path("/orderBooks/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Order Book of Specified Pair in Bitfinex", responseContainer = "List", response = BitfinexOrderBook.class))
    public Response getOrderBookForPair(@PathParam("pair") @Valid @NotNull BitfinexPair pair) throws Exception {
        return Response.ok(marketService.getExchangeOrderBookForPair(pair)).build();
    }
}

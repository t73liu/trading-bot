package io.github.t73liu.rest;

import eu.verdelhan.ta4j.Tick;
import io.github.t73liu.exception.ExceptionWrapper;
import io.github.t73liu.exchange.poloniex.rest.PoloniexAccountService;
import io.github.t73liu.exchange.poloniex.rest.PoloniexMarketService;
import io.github.t73liu.exchange.poloniex.rest.PoloniexOrderService;
import io.github.t73liu.model.poloniex.PoloniexCandleInterval;
import io.github.t73liu.model.poloniex.PoloniexOrderBook;
import io.github.t73liu.model.poloniex.PoloniexPair;
import io.github.t73liu.model.poloniex.PoloniexTicker;
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
import java.math.BigDecimal;
import java.util.Map;

@Component
@Path("/poloniex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("PoloniexResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class PoloniexResource {
    private final PoloniexAccountService accountService;
    private final PoloniexMarketService marketService;
    private final PoloniexOrderService orderService;

    @Autowired
    public PoloniexResource(PoloniexAccountService service, PoloniexMarketService marketService, PoloniexOrderService orderService) {
        this.accountService = service;
        this.marketService = marketService;
        this.orderService = orderService;
    }

    @GET
    @Path("/candles/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Checks if there is candlestick opportunity", responseContainer = "List", response = Tick.class))
    public Response checkCandles(@PathParam("pair") @Valid @NotNull PoloniexPair pair,
                                 @QueryParam("interval") @Valid @NotNull PoloniexCandleInterval interval,
                                 @QueryParam("startSeconds") long startSeconds,
                                 @QueryParam("endSeconds") long endSeconds) throws Exception {
        return Response.ok(marketService.getCandlestickForPair(pair, startSeconds, endSeconds, interval)).build();
    }

    @GET
    @Path("/tickers")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved All Tickers in Poloniex", response = Map.class))
    public Response getAllTickers() throws Exception {
        return Response.ok(marketService.getAllTicker()).build();
    }

    @GET
    @Path("/tickers/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Ticker of Specified Pair in Poloniex", response = PoloniexTicker.class))
    public Response getTickerForPair(@PathParam("pair") @Valid @NotNull PoloniexPair pair) throws Exception {
        return Response.ok(marketService.getTickerForPair(pair)).build();
    }

    @GET
    @Path("/orderBooks")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved All Order Books in Poloniex", response = Map.class))
    public Response getAllOrderBooks(@QueryParam("depth") @DefaultValue("3") int depth) throws Exception {
        return Response.ok(marketService.getAllOrderBook(depth)).build();
    }

    @GET
    @Path("/orderBooks/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Order Book of Specified Pair in Poloniex", response = PoloniexOrderBook.class))
    public Response getOrderBookForPair(@PathParam("pair") @Valid @NotNull PoloniexPair pair,
                                        @QueryParam("depth") @DefaultValue("3") int depth) throws Exception {
        return Response.ok(marketService.getOrderBookForPair(pair, depth)).build();
    }

    @GET
    @Path("/balances")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Balances in Poloniex", response = Map.class))
    public Response getBalances() throws Exception {
        return Response.ok(accountService.getCompleteBalances()).build();
    }

    @GET
    @Path("/orders/open")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Open Orders in Poloniex", response = Map.class))
    public Response getOpenOrders(@QueryParam("pair") PoloniexPair pair) throws Exception {
        return Response.ok(orderService.getOpenOrders(pair)).build();
    }

    @GET
    @Path("/orders/buy")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Open Orders in Poloniex", response = Map.class))
    public Response placeBuy(@QueryParam("pair") @Valid @NotNull PoloniexPair pair,
                             @QueryParam("orderType") @DefaultValue("buy") String orderType,
                             @QueryParam("fulfillmentType") @DefaultValue("immediateOrCancel") String fulfillmentType) throws Exception {
        return Response.ok(orderService.placeOrder(pair, BigDecimal.valueOf(0), BigDecimal.valueOf(0), orderType, fulfillmentType)).build();
    }
}

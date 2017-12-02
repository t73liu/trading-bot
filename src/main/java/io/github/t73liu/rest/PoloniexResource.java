package io.github.t73liu.rest;

import io.github.t73liu.exchange.poloniex.rest.PoloniexAccountService;
import io.github.t73liu.exchange.poloniex.rest.PoloniexMarketService;
import io.github.t73liu.exchange.poloniex.rest.PoloniexOrderService;
import io.github.t73liu.model.poloniex.PoloniexCandleInterval;
import io.github.t73liu.model.poloniex.PoloniexOrderBook;
import io.github.t73liu.model.poloniex.PoloniexPair;
import io.github.t73liu.model.poloniex.PoloniexTicker;
import io.swagger.v3.oas.annotations.media.ArraySchema;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;
import io.swagger.v3.oas.annotations.tags.Tag;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.ta4j.core.BaseTick;

import javax.validation.Valid;
import javax.validation.constraints.NotNull;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.math.BigDecimal;

@Component
@Path("/poloniex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Tag(name = "PoloniexResource")
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
    @ApiResponse(responseCode = "200", description = "Retrieved Candle Stick of Specified Pair in Poloniex", content = @Content(array = @ArraySchema(schema = @Schema(implementation = BaseTick.class))))
    public Response checkCandles(@PathParam("pair") @Valid @NotNull PoloniexPair pair,
                                 @QueryParam("interval") @Valid @NotNull PoloniexCandleInterval interval,
                                 @QueryParam("startSeconds") long startSeconds,
                                 @QueryParam("endSeconds") long endSeconds) throws Exception {
        return Response.ok(marketService.getCandlestickForPair(pair, startSeconds, endSeconds, interval)).build();
    }

    @GET
    @Path("/tickers")
    @ApiResponse(responseCode = "200", description = "Retrieved All Tickers in Poloniex", content = @Content(schema = @Schema(implementation = Object.class)))
    public Response getAllTickers() throws Exception {
        return Response.ok(marketService.getAllTicker()).build();
    }

    @GET
    @Path("/tickers/{pair}")
    @ApiResponse(responseCode = "200", description = "Retrieved Ticker of Specified Pair in Poloniex", content = @Content(schema = @Schema(implementation = PoloniexTicker.class)))
    public Response getTickerForPair(@PathParam("pair") @Valid @NotNull PoloniexPair pair) throws Exception {
        return Response.ok(marketService.getTickerForPair(pair)).build();
    }

    @GET
    @Path("/orderBooks")
    @ApiResponse(responseCode = "200", description = "Retrieved All Order Books in Poloniex", content = @Content(schema = @Schema(implementation = Object.class)))
    public Response getAllOrderBooks(@QueryParam("depth") @DefaultValue("3") int depth) throws Exception {
        return Response.ok(marketService.getAllOrderBook(depth)).build();
    }

    @GET
    @Path("/orderBooks/{pair}")
    @ApiResponse(responseCode = "200", description = "Retrieved Order Book of Specified Pair in Poloniex", content = @Content(schema = @Schema(implementation = PoloniexOrderBook.class)))
    public Response getOrderBookForPair(@PathParam("pair") @Valid @NotNull PoloniexPair pair,
                                        @QueryParam("depth") @DefaultValue("3") int depth) throws Exception {
        return Response.ok(marketService.getOrderBookForPair(pair, depth)).build();
    }

    @GET
    @Path("/balances")
    @ApiResponse(responseCode = "200", description = "Retrieved Balances in Poloniex", content = @Content(schema = @Schema(implementation = Object.class)))
    public Response getBalances() throws Exception {
        return Response.ok(accountService.getCompleteBalances()).build();
    }

    @GET
    @Path("/orders/open")
    @ApiResponse(responseCode = "200", description = "Retrieved Open Orders in Poloniex", content = @Content(schema = @Schema(implementation = Object.class)))
    public Response getOpenOrders(@QueryParam("pair") PoloniexPair pair) throws Exception {
        return Response.ok(orderService.getOpenOrders(pair)).build();
    }

    @GET
    @Path("/orders/buy")
    @ApiResponse(responseCode = "200", description = "Place Buy Order in Poloniex", content = @Content(schema = @Schema(implementation = Object.class)))
    public Response placeBuy(@QueryParam("pair") @Valid @NotNull PoloniexPair pair,
                             @QueryParam("orderType") @DefaultValue("buy") String orderType,
                             @QueryParam("fulfillmentType") @DefaultValue("immediateOrCancel") String fulfillmentType) throws Exception {
        return Response.ok(orderService.placeOrder(pair, BigDecimal.valueOf(0), BigDecimal.valueOf(0), orderType, fulfillmentType)).build();
    }
}

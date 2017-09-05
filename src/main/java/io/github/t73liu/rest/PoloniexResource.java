package io.github.t73liu.rest;

import io.github.t73liu.model.ExceptionWrapper;
import io.github.t73liu.model.currency.PoloniexPair;
import io.github.t73liu.service.PoloniexService;
import io.github.t73liu.service.PoloniexTicker;
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
import java.util.Map;

@Component
@Path("/poloniex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("PoloniexResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class PoloniexResource {
    private final PoloniexService service;
    private final PoloniexTicker ticker;

    @Autowired
    public PoloniexResource(PoloniexService service, PoloniexTicker ticker) {
        this.service = service;
        this.ticker = ticker;
    }

    @GET
    @Path("/arbitrage")
    @ApiResponses(@ApiResponse(code = 200, message = "Checks if there is arbitrage opportunity", response = Boolean.class))
    public Response checkArbitrage() throws Exception {
        return Response.ok(service.checkArbitrage()).build();
    }

    @GET
    @Path("/tickers")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Ticker of Specified Pair in Poloniex", response = Map.class))
    public Response getTicker(@QueryParam("pair") @Valid @NotNull PoloniexPair pair) throws Exception {
        return Response.ok(ticker.getTickerValue(pair)).build();
    }

    @GET
    @Path("/balances")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Balances in Poloniex", response = Map.class))
    public Response getBalances() throws Exception {
        return Response.ok(service.getCompleteBalances()).build();
    }

    @GET
    @Path("/orders/open")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Open Orders in Poloniex", response = Map.class))
    public Response getOpenOrders(@QueryParam("pair") PoloniexPair pair) throws Exception {
        return Response.ok(service.getOpenOrders(pair)).build();
    }

    @GET
    @Path("/orders/buy")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Open Orders in Poloniex", response = Map.class))
    public Response placeBuy(@QueryParam("pair") @Valid @NotNull PoloniexPair pair,
                             @QueryParam("orderType") @DefaultValue("buy") String orderType,
                             @QueryParam("fulfillmentType") @DefaultValue("immediateOrCancel") String fulfillmentType) throws Exception {
        return Response.ok(service.placeOrder(pair, 0, 0, orderType, fulfillmentType)).build();
    }
}

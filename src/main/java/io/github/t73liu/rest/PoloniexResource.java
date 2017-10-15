package io.github.t73liu.rest;

import eu.verdelhan.ta4j.Tick;
import io.github.t73liu.exception.ExceptionWrapper;
import io.github.t73liu.exchange.poloniex.rest.PoloniexService;
import io.github.t73liu.model.currency.PoloniexPair;
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
import java.time.LocalDateTime;
import java.util.Map;

import static io.github.t73liu.model.CandlestickIntervals.THIRTY_MIN;
import static io.github.t73liu.util.DateUtil.getCurrentLocalDateTime;

@Component
@Path("/poloniex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("PoloniexResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class PoloniexResource {
    private final PoloniexService service;

    @Autowired
    public PoloniexResource(PoloniexService service) {
        this.service = service;
    }

    @GET
    @Path("/arbitrage")
    @ApiResponses(@ApiResponse(code = 200, message = "Checks if there is arbitrage opportunity", response = Boolean.class))
    public Response checkArbitrage() throws Exception {
        return Response.ok(service.checkArbitrage()).build();
    }

    @GET
    @Path("/candles")
    @ApiResponses(@ApiResponse(code = 200, message = "Checks if there is candlestick opportunity", responseContainer = "List", response = Tick.class))
    public Response checkCandles() throws Exception {
        LocalDateTime endLocalDateTime = getCurrentLocalDateTime();
        LocalDateTime startLocalDateTime = endLocalDateTime.minusHours(6);
        return Response.ok(service.getCandlestick(PoloniexPair.USDT_XRP, startLocalDateTime, endLocalDateTime, THIRTY_MIN)).build();
    }

    @GET
    @Path("/tickers")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Ticker of Specified Pair in Poloniex", response = Map.class))
    public Response getTicker(@QueryParam("pair") @Valid @NotNull PoloniexPair pair) throws Exception {
        return Response.ok(service.getTickerValue(pair)).build();
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
        return Response.ok(service.placeOrder(pair, BigDecimal.valueOf(0), BigDecimal.valueOf(0), orderType, fulfillmentType)).build();
    }
}

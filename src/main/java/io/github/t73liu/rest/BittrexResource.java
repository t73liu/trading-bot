package io.github.t73liu.rest;

import io.github.t73liu.exception.ExceptionWrapper;
import io.github.t73liu.exchange.bittrex.rest.BittrexAccountService;
import io.github.t73liu.exchange.bittrex.rest.BittrexMarketService;
import io.github.t73liu.exchange.bittrex.rest.BittrexOrderService;
import io.github.t73liu.model.bittrex.BittrexOrderBook;
import io.github.t73liu.model.bittrex.BittrexPair;
import io.github.t73liu.model.bittrex.BittrexTicker;
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
@Path("/bittrex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("BittrexResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class BittrexResource {
    private final BittrexAccountService accountService;
    private final BittrexMarketService marketService;
    private final BittrexOrderService orderService;

    @Autowired
    public BittrexResource(BittrexAccountService accountService, BittrexMarketService marketService, BittrexOrderService orderService) {
        this.accountService = accountService;
        this.marketService = marketService;
        this.orderService = orderService;
    }

    @GET
    @Path("/tickers/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Ticker of Specified Pair in Bittrex", response = BittrexTicker.class))
    public Response getTickers(@PathParam("pair") @Valid @NotNull BittrexPair pair) throws Exception {
        return Response.ok(marketService.getTickerForPair(pair)).build();
    }

    @GET
    @Path("/orderBooks/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Order Book of Specified Pair in Bittrex", response = BittrexOrderBook.class))
    public Response getOrderBookForPair(@PathParam("pair") @Valid @NotNull BittrexPair pair) throws Exception {
        return Response.ok(marketService.getOrderBookForPair(pair)).build();
    }
}

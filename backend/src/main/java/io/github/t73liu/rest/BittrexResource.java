package io.github.t73liu.rest;

import io.github.t73liu.exchange.bittrex.rest.BittrexAccountService;
import io.github.t73liu.exchange.bittrex.rest.BittrexMarketService;
import io.github.t73liu.exchange.bittrex.rest.BittrexOrderService;
import io.github.t73liu.model.bittrex.BittrexOrderBook;
import io.github.t73liu.model.bittrex.BittrexPair;
import io.github.t73liu.model.bittrex.BittrexTicker;
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
@Path("/bittrex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Tag(name = "BittrexResource")
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
    @ApiResponse(responseCode = "200", description = "Retrieved Ticker of Specified Pair in Bittrex", content = @Content(schema = @Schema(implementation = BittrexTicker.class)))
    public Response getTickers(@PathParam("pair") @Valid @NotNull BittrexPair pair) throws Exception {
        return Response.ok(marketService.getTickerForPair(pair)).build();
    }

    @GET
    @Path("/orderBooks/{pair}")
    @ApiResponse(responseCode = "200", description = "Retrieved Order Book of Specified Pair in Bittrex", content = @Content(schema = @Schema(implementation = BittrexOrderBook.class)))
    public Response getOrderBookForPair(@PathParam("pair") @Valid @NotNull BittrexPair pair) throws Exception {
        return Response.ok(marketService.getOrderBookForPair(pair)).build();
    }
}

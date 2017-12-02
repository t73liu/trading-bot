package io.github.t73liu.rest;

import io.github.t73liu.exchange.quadriga.rest.QuadrigaAccountService;
import io.github.t73liu.exchange.quadriga.rest.QuadrigaMarketService;
import io.github.t73liu.exchange.quadriga.rest.QuadrigaOrderService;
import io.github.t73liu.model.quadriga.QuadrigaOrderBook;
import io.github.t73liu.model.quadriga.QuadrigaPair;
import io.github.t73liu.model.quadriga.QuadrigaTicker;
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
@Path("/quadriga")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Tag(name = "QuadrigaResource")
public class QuadrigaResource {
    private final QuadrigaAccountService accountService;
    private final QuadrigaMarketService marketService;
    private final QuadrigaOrderService orderService;

    @Autowired
    public QuadrigaResource(QuadrigaAccountService accountService, QuadrigaMarketService marketService, QuadrigaOrderService orderService) {
        this.accountService = accountService;
        this.marketService = marketService;
        this.orderService = orderService;
    }

    @GET
    @Path("/tickers/{pair}")
    @ApiResponse(responseCode = "200", description = "Retrieved Ticker of Specified Pair in Quadriga", content = @Content(schema = @Schema(implementation = QuadrigaTicker.class)))
    public Response getTickerForPair(@PathParam("pair") @Valid @NotNull QuadrigaPair pair) throws Exception {
        return Response.ok(marketService.getTickerForPair(pair)).build();
    }

    @GET
    @Path("/orderBooks/{pair}")
    @ApiResponse(responseCode = "200", description = "Retrieved Order Book of Specified Pair in Quadriga", content = @Content(schema = @Schema(implementation = QuadrigaOrderBook.class)))
    public Response getOrderBookForPair(@PathParam("pair") @Valid @NotNull QuadrigaPair pair) throws Exception {
        return Response.ok(marketService.getOrderBookForPair(pair)).build();
    }
}

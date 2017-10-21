package io.github.t73liu.rest;

import io.github.t73liu.exception.ExceptionWrapper;
import io.github.t73liu.exchange.quadriga.rest.QuadrigaAccountService;
import io.github.t73liu.exchange.quadriga.rest.QuadrigaMarketService;
import io.github.t73liu.exchange.quadriga.rest.QuadrigaOrderService;
import io.github.t73liu.model.quadriga.QuadrigaOrderBook;
import io.github.t73liu.model.quadriga.QuadrigaPair;
import io.github.t73liu.model.quadriga.QuadrigaTicker;
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
@Path("/quadriga")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("QuadrigaResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
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
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Ticker of Specified Pair in Quadriga", response = QuadrigaTicker.class))
    public Response getTickerForPair(@PathParam("pair") @Valid @NotNull QuadrigaPair pair) throws Exception {
        return Response.ok(marketService.getTickerForPair(pair)).build();
    }

    @GET
    @Path("/orderBooks/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Order Book of Specified Pair in Quadriga", response = QuadrigaOrderBook.class))
    public Response getOrderBookForPair(@PathParam("pair") @Valid @NotNull QuadrigaPair pair) throws Exception {
        return Response.ok(marketService.getOrderBookForPair(pair)).build();
    }
}

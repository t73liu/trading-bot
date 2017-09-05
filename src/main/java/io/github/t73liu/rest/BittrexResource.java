package io.github.t73liu.rest;

import io.github.t73liu.model.ExceptionWrapper;
import io.github.t73liu.service.BittrexService;
import io.github.t73liu.service.BittrexTicker;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiResponse;
import io.swagger.annotations.ApiResponses;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.ws.rs.Consumes;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Map;

@Component
@Path("/bittrex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("BittrexResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class BittrexResource {
    private final BittrexService service;
    private final BittrexTicker ticker;

    @Autowired
    public BittrexResource(BittrexService service, BittrexTicker ticker) {
        this.service = service;
        this.ticker = ticker;
    }

    @GET
    @Path("/tickers")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Currency Information", responseContainer = "List", response = Map.class))
    public Response getTickers() {
        return Response.ok().build();
    }
}

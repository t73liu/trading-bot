package io.github.t73liu.rest;

import io.github.t73liu.exception.ExceptionWrapper;
import io.github.t73liu.exchange.bitfinex.rest.BitfinexService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiResponse;
import io.swagger.annotations.ApiResponses;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.validation.constraints.NotNull;
import javax.ws.rs.*;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.util.Map;

@Component
@Path("/bitfinex")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("BitfinexResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class BitfinexResource {
    private final BitfinexService service;

    @Autowired
    public BitfinexResource(BitfinexService service) {
        this.service = service;
    }

    @GET
    @Path("/ticker/{pair}")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Currency Information", response = Map.class))
    public Response getTickers(@PathParam("pair") @NotNull String pair) {
        return Response.ok().build();
    }
}

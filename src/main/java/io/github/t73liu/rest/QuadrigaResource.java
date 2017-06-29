package io.github.t73liu.rest;

import io.github.t73liu.model.ExceptionWrapper;
import io.github.t73liu.model.Order;
import io.github.t73liu.util.DateUtil;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiParam;
import io.swagger.annotations.ApiResponse;
import io.swagger.annotations.ApiResponses;
import org.glassfish.jersey.client.ClientConfig;
import org.glassfish.jersey.jackson.JacksonFeature;
import org.springframework.stereotype.Component;

import javax.ws.rs.*;
import javax.ws.rs.client.Client;
import javax.ws.rs.client.ClientBuilder;
import javax.ws.rs.core.MediaType;
import javax.ws.rs.core.Response;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@Component
@Path("/quadriga")
@Consumes(MediaType.APPLICATION_JSON)
@Produces(MediaType.APPLICATION_JSON)
@Api("QuadrigaResource")
@ApiResponses(@ApiResponse(code = 500, message = "Internal Server Error", response = ExceptionWrapper.class))
public class QuadrigaResource {
    @GET
    @Path("/orders")
    @ApiResponses(@ApiResponse(code = 200, message = "Retrieved Orders of selected Currencies", responseContainer = "List", response = Order.class))
    public Response getOrders(@ApiParam(example = "btc_cad,btc_usd,eth_btc,eth_cad", required = true) @QueryParam("currencyPair") String currencyPair) {
        ClientConfig cc = new ClientConfig().register(new JacksonFeature());
        Client client = ClientBuilder.newClient(cc);
        Map<String, Object> response = client.target("https://api.quadrigacx.com/v2/order_book?book=" + currencyPair).request().get().readEntity(Map.class);
        List<Order> output = new ArrayList<>();
        LocalDateTime triggerTime = DateUtil.convertUnixTimestamp(response.get("timestamp").toString());
        ((List<List<String>>) response.get("bids")).forEach(priceQuantityList -> {
            Order order = new Order();
            order.setIssueTime(triggerTime);
            order.setType("BID");
            order.setQuantity(Double.parseDouble(priceQuantityList.get(1)));
            order.setPrice(Double.parseDouble(priceQuantityList.get(0)));
            output.add(order);
        });
        ((List<List<String>>) response.get("asks")).forEach(priceQuantityList -> {
            Order order = new Order();
            order.setIssueTime(triggerTime);
            order.setType("ASK");
            order.setQuantity(Double.parseDouble(priceQuantityList.get(1)));
            order.setPrice(Double.parseDouble(priceQuantityList.get(0)));
            output.add(order);
        });
        return Response.ok(output).build();
    }
}

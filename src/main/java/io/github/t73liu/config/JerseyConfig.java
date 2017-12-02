package io.github.t73liu.config;

import com.fasterxml.jackson.jaxrs.json.JacksonJaxbJsonProvider;
import com.google.common.collect.ImmutableSet;
import io.github.t73liu.provider.*;
import io.swagger.v3.jaxrs2.SwaggerSerializers;
import io.swagger.v3.jaxrs2.integration.JaxrsOpenApiContextBuilder;
import io.swagger.v3.jaxrs2.integration.resources.OpenApiResource;
import io.swagger.v3.oas.integration.SwaggerConfiguration;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Contact;
import io.swagger.v3.oas.models.info.Info;
import io.swagger.v3.oas.models.info.License;
import org.glassfish.jersey.jackson.JacksonFeature;
import org.glassfish.jersey.server.ResourceConfig;
import org.glassfish.jersey.server.ServerProperties;
import org.glassfish.jersey.server.wadl.internal.WadlResource;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.Configuration;

import javax.annotation.PostConstruct;
import javax.ws.rs.Path;

@Configuration
public class JerseyConfig extends ResourceConfig {
    private static final Logger LOGGER = LoggerFactory.getLogger(JerseyConfig.class);

    @Value("${app.version:1.0.0-SNAPSHOT}")
    private String appVersion;

    @Value("${spring.jersey.application-path:/}")
    private String apiPath;

    private ApplicationContext context;

    @Autowired
    public JerseyConfig(ApplicationContext context) {
        this.context = context;
        configureProperties();
        setupResources();
        registerProviders();
    }

    private void setupResources() {
        context.getBeansWithAnnotation(Path.class).forEach((name, resource) -> {
            LOGGER.info("Registering Jersey Resource: {}", name);
            register(resource);
        });
    }

    private void registerProviders() {
        // General Providers
        register(JacksonFeature.class);
        register(JacksonJaxbJsonProvider.class);

        // Internal Custom Providers
        register(ObjectMapperContextResolver.class);
        register(LocalDateParamProvider.class);
        register(GeneralExceptionMapper.class);
        register(JsonProcessingExceptionMapper.class);
        register(ValidationExceptionMapper.class);

        // Swagger Providers
        register(OpenApiResource.class);
        register(SwaggerSerializers.class);
        register(WadlResource.class);
    }

    private void configureProperties() {
        property(ServerProperties.BV_SEND_ERROR_IN_RESPONSE, Boolean.TRUE);
    }

    @PostConstruct
    private void initializeSwagger() throws Exception {
        OpenAPI oas = new OpenAPI();
        Info info = new Info()
                .title("Crypto-Currency Trading Bot")
                .description("This is a crypto-currency trading bot with a front end and custom news feeds.")
                .version(this.appVersion)
                .contact(new Contact()
                        .email("cryptotest92@gmail.com"))
                .license(new License()
                        .name("Apache 2.0")
                        .url("http://www.apache.org/licenses/LICENSE-2.0.html"));
        oas.info(info);
        SwaggerConfiguration oasConfig = new SwaggerConfiguration()
                .prettyPrint(true)
                .openAPI(oas)
                .resourcePackages(ImmutableSet.of("io.github.t73liu.rest"));
        new JaxrsOpenApiContextBuilder()
                .openApiConfiguration(oasConfig)
                .buildContext(true);
    }
}

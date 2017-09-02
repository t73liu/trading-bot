package io.github.t73liu.config;

import com.fasterxml.jackson.jaxrs.json.JacksonJaxbJsonProvider;
import io.github.t73liu.provider.GeneralExceptionMapper;
import io.github.t73liu.provider.LocalDateParamProvider;
import io.github.t73liu.provider.ObjectMapperContextResolver;
import io.swagger.jaxrs.config.BeanConfig;
import io.swagger.jaxrs.listing.ApiListingResource;
import io.swagger.jaxrs.listing.SwaggerSerializers;
import io.swagger.models.Contact;
import io.swagger.models.Info;
import io.swagger.models.License;
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
    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());

    @Value("${app.version:1.0.0-SNAPSHOT}")
    private String appVersion;

    @Value("${spring.jersey.application-path:/}")
    private String apiPath;

    private ApplicationContext context;

    @Autowired
    public JerseyConfig(ApplicationContext context) {
        this.context = context;
        setupResources();
        registerProviders();
        configureProperties();
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
        // FIXME implement CORS?
//        register(CORSFilter.class);

        // Internal Custom Providers
        register(ObjectMapperContextResolver.class);
        register(LocalDateParamProvider.class);
        register(GeneralExceptionMapper.class);

        // Swagger Providers
        register(ApiListingResource.class);
        register(SwaggerSerializers.class);
        register(WadlResource.class);
    }

    private void configureProperties() {
        property(ServerProperties.BV_SEND_ERROR_IN_RESPONSE, Boolean.TRUE);
    }

    @PostConstruct
    private void initializeSwagger() {
        BeanConfig beanConfig = new BeanConfig();
        beanConfig.setResourcePackage("io.github.t73liu.rest");
        beanConfig.setScan(true);
        beanConfig.setPrettyPrint(true);
        beanConfig.setBasePath(this.apiPath);

        Info info = new Info();
        beanConfig.setInfo(info);
        info.setTitle("Crypto-Currency Trading Bot");
        info.setVersion(this.appVersion);

        Contact contact = new Contact();
        info.setContact(contact);

        License license = new License();
        info.setLicense(license);
    }
}

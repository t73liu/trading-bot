package io.github.t73liu.util;

import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateDeserializer;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateSerializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;
import com.fasterxml.jackson.module.afterburner.AfterburnerModule;

import java.time.LocalDate;
import java.time.LocalDateTime;

public class ObjectMapperFactory {
    public static final ObjectMapper OBJECT_MAPPER = createObjectMapper();

    public static final ObjectReader OBJECT_READER = OBJECT_MAPPER.reader();

    public static final ObjectWriter OBJECT_WRITER = OBJECT_MAPPER.writer();

    private static ObjectMapper createObjectMapper() {
        ObjectMapper mapper = new ObjectMapper();
        mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        mapper.configure(MapperFeature.ACCEPT_CASE_INSENSITIVE_PROPERTIES, true);
        mapper.configure(SerializationFeature.FAIL_ON_EMPTY_BEANS, false);
//        mapper.configure(SerializationFeature.INDENT_OUTPUT, true);
        JavaTimeModule timeModule = new JavaTimeModule();
        timeModule.addDeserializer(LocalDate.class, new LocalDateDeserializer(DateUtil.LOCALDATE_ISO_FORMATTER));
        timeModule.addDeserializer(LocalDateTime.class, new LocalDateTimeDeserializer(DateUtil.LOCALDATETIME_ISO_FORMATTER));
        timeModule.addSerializer(LocalDate.class, new LocalDateSerializer(DateUtil.LOCALDATE_ISO_FORMATTER));
        timeModule.addSerializer(LocalDateTime.class, new LocalDateTimeSerializer(DateUtil.LOCALDATETIME_ISO_FORMATTER));
        mapper.registerModule(timeModule);
        mapper.registerModule(new AfterburnerModule());
        return mapper;
    }
}

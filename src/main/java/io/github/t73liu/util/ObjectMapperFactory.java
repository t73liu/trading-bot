package io.github.t73liu.util;

import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.dataformat.csv.CsvMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateDeserializer;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateSerializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;
import com.fasterxml.jackson.module.afterburner.AfterburnerModule;

import java.time.LocalDate;
import java.time.LocalDateTime;

public class ObjectMapperFactory {
    public static final ObjectMapper JSON_MAPPER = createJSONMapper();

    public static final ObjectReader JSON_READER = JSON_MAPPER.reader();

    public static final ObjectWriter JSON_WRITER = JSON_MAPPER.writer();

    public static final CsvMapper CSV_MAPPER = createCSVMapper();

    private static ObjectMapper createJSONMapper() {
        ObjectMapper mapper = new ObjectMapper();
        configureMapper(mapper);
        return mapper;
    }

    private static CsvMapper createCSVMapper() {
        CsvMapper mapper = new CsvMapper();
        mapper.disable(MapperFeature.SORT_PROPERTIES_ALPHABETICALLY);
        configureMapper(mapper);
        return mapper;
    }

    private static void configureMapper(ObjectMapper mapper) {
        mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        mapper.configure(MapperFeature.ACCEPT_CASE_INSENSITIVE_PROPERTIES, true);
        mapper.configure(SerializationFeature.FAIL_ON_EMPTY_BEANS, false);
        JavaTimeModule timeModule = new JavaTimeModule();
        timeModule.addDeserializer(LocalDate.class, new LocalDateDeserializer(DateUtil.LOCALDATE_ISO_FORMATTER));
        timeModule.addDeserializer(LocalDateTime.class, new LocalDateTimeDeserializer(DateUtil.LOCALDATETIME_ISO_FORMATTER));
        timeModule.addSerializer(LocalDate.class, new LocalDateSerializer(DateUtil.LOCALDATE_ISO_FORMATTER));
        timeModule.addSerializer(LocalDateTime.class, new LocalDateTimeSerializer(DateUtil.LOCALDATETIME_ISO_FORMATTER));
        mapper.registerModule(timeModule);
        mapper.registerModule(new AfterburnerModule());
    }
}

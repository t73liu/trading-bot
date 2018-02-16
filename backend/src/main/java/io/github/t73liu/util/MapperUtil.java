package io.github.t73liu.util;

import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.databind.type.TypeFactory;
import com.fasterxml.jackson.dataformat.csv.CsvMapper;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateDeserializer;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateSerializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;
import com.fasterxml.jackson.datatype.jsr310.ser.ZonedDateTimeSerializer;
import com.fasterxml.jackson.module.afterburner.AfterburnerModule;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.ZonedDateTime;
import java.util.List;

public class MapperUtil {
    public static final ObjectMapper JSON_MAPPER = createJSONMapper();

    public static final ObjectReader JSON_READER = JSON_MAPPER.reader();

    public static final ObjectWriter JSON_WRITER = JSON_MAPPER.writer();

    public static final TypeFactory TYPE_FACTORY = JSON_MAPPER.getTypeFactory();

    public static final CsvMapper CSV_MAPPER = createCSVMapper();

    public static final ObjectReader CSV_READER = CSV_MAPPER.reader();

    public static final ObjectWriter CSV_WRITER = CSV_MAPPER.writer();

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
        timeModule.addSerializer(ZonedDateTime.class, new ZonedDateTimeSerializer(DateUtil.ZONED_DATETIME_ISO_FORMATTER));
        mapper.registerModule(timeModule);
        mapper.registerModule(new AfterburnerModule());
    }

    public static void writeCSV(Object data, Class dataClass, String path) throws Exception {
        CSV_WRITER.with(CSV_MAPPER.schemaFor(dataClass).withHeader())
                .writeValue(Files.newOutputStream(Paths.get(path)), data);
    }

    public static <Type> List<Type> readCSV(Class<Type> dataClass, String path) throws Exception {
        MappingIterator<Type> iterator = CSV_READER.forType(dataClass)
                .with(CSV_MAPPER.schemaFor(dataClass).withHeader())
                .readValues(Files.newInputStream(Paths.get(path)));
        return iterator.readAll();
    }
}

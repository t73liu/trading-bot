package io.github.t73liu.util;

import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.MapperFeature;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.SerializationFeature;
import com.fasterxml.jackson.databind.type.CollectionType;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateDeserializer;
import com.fasterxml.jackson.datatype.jsr310.deser.LocalDateTimeDeserializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateSerializer;
import com.fasterxml.jackson.datatype.jsr310.ser.LocalDateTimeSerializer;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.util.List;

public class ObjectMapperFactory {
    public static ObjectMapper getNewInstance() {
        ObjectMapper mapper = new ObjectMapper();
        mapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        mapper.configure(MapperFeature.ACCEPT_CASE_INSENSITIVE_PROPERTIES, true);
        mapper.configure(SerializationFeature.FAIL_ON_EMPTY_BEANS, false);
        mapper.configure(SerializationFeature.INDENT_OUTPUT, true);
        JavaTimeModule timeModule = new JavaTimeModule();
        timeModule.addDeserializer(LocalDate.class, new LocalDateDeserializer(DateUtil.LOCALDATE_ISO_FORMATTER));
        timeModule.addDeserializer(LocalDateTime.class, new LocalDateTimeDeserializer(DateUtil.LOCALDATETIME_ISO_FORMATTER));
        timeModule.addSerializer(LocalDate.class, new LocalDateSerializer(DateUtil.LOCALDATE_ISO_FORMATTER));
        timeModule.addSerializer(LocalDateTime.class, new LocalDateTimeSerializer(DateUtil.LOCALDATETIME_ISO_FORMATTER));
        mapper.registerModule(timeModule);
        return mapper;
    }

    public static CollectionType getList(ObjectMapper mapper, Class dataClass) {
        return mapper.getTypeFactory().constructCollectionType(List.class, dataClass);
    }
}

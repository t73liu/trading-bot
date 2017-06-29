package io.github.t73liu.provider;

import io.github.t73liu.util.DateUtil;

import javax.annotation.PostConstruct;
import javax.ws.rs.ext.ParamConverter;
import javax.ws.rs.ext.ParamConverterProvider;
import javax.ws.rs.ext.Provider;
import java.lang.annotation.Annotation;
import java.lang.reflect.Type;
import java.time.LocalDate;

import static org.apache.commons.lang3.StringUtils.isBlank;

@Provider
public class LocalDateParamProvider implements ParamConverterProvider {
    private ParamConverter<LocalDate> localDateParamConverter;

    @PostConstruct
    private void initializeLocalDateConverter() {
        this.localDateParamConverter = new ParamConverter<LocalDate>() {
            @Override
            public LocalDate fromString(String dateStr) {
                try {
                    if (isBlank(dateStr)) {
                        return DateUtil.getCurrentLocalDate();
                    }
                    return dateStr.length() == 10 ? DateUtil.parseLocalDateISO(dateStr) : DateUtil.parseLocalDateShort(dateStr);
                } catch (Exception ex) {
                    throw new IllegalArgumentException("Please provide date string with the following format yyyyMMdd or yyyy-MM-dd");
                }
            }

            @Override
            public String toString(LocalDate localDate) {
                return DateUtil.formatLocalDateISO(localDate);
            }
        };
    }

    @Override
    public <T> ParamConverter<T> getConverter(Class<T> aClass, Type type, Annotation[] annotations) {
        return LocalDate.class == aClass ? (ParamConverter<T>) localDateParamConverter : null;
    }
}

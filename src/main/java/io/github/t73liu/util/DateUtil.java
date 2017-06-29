package io.github.t73liu.util;

import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.time.format.DateTimeFormatter;

public class DateUtil {
    public static final DateTimeFormatter LOCALDATE_ISO_FORMATTER = DateTimeFormatter.ISO_LOCAL_DATE;
    public static final DateTimeFormatter LOCALDATE_SHORT_FORMATTER = DateTimeFormatter.ofPattern("uuuuMMdd");
    public static final String TIMEZONE = "America/New_York";

    public static LocalDateTime convertUnixTimestamp(long timestampInSeconds) {
        return Instant.ofEpochSecond(timestampInSeconds)
                .atZone(ZoneId.of(TIMEZONE))
                .toLocalDateTime();
    }

    public static LocalDateTime convertUnixTimestamp(String timestampInSeconds) {
        return convertUnixTimestamp(Long.parseLong(timestampInSeconds));
    }

    public static LocalDate getCurrentLocalDate() {
        return LocalDate.now(ZoneId.of(TIMEZONE));
    }

    public static LocalDate parseLocalDateISO(String dateStr) {
        return LocalDate.parse(dateStr, LOCALDATE_ISO_FORMATTER);
    }

    public static LocalDate parseLocalDateShort(String dateStr) {
        return LocalDate.parse(dateStr, LOCALDATE_SHORT_FORMATTER);
    }

    public static String formatLocalDateISO(LocalDate localDate) {
        return localDate.format(LOCALDATE_ISO_FORMATTER);
    }

    public static String formatLocalDateShort(LocalDate localDate) {
        return localDate.format(LOCALDATE_SHORT_FORMATTER);
    }
}

package io.github.t73liu.util;

import java.time.*;
import java.time.format.DateTimeFormatter;

public class DateUtil {
    public static final DateTimeFormatter LOCALDATE_ISO_FORMATTER = DateTimeFormatter.ISO_LOCAL_DATE;
    public static final DateTimeFormatter LOCALDATE_SHORT_FORMATTER = DateTimeFormatter.ofPattern("uuuuMMdd");
    public static final DateTimeFormatter LOCALDATETIME_ISO_FORMATTER = DateTimeFormatter.ISO_LOCAL_DATE_TIME;
    public static final String TIMEZONE = "America/New_York";
    public static final ZoneId TIMEZONE_ID = ZoneId.of(TIMEZONE);

    public static LocalDate getCurrentLocalDate() {
        return LocalDate.now(ZoneId.of(TIMEZONE));
    }

    public static LocalDateTime getCurrentLocalDateTime() {
        return LocalDateTime.now(ZoneId.of(TIMEZONE));
    }

    public static LocalDateTime unixTimeStampToLocalDateTime(long timestampInSeconds) {
        return Instant.ofEpochSecond(timestampInSeconds)
                .atZone(TIMEZONE_ID)
                .toLocalDateTime();
    }

    public static LocalDateTime unixTimeStampToLocalDateTime(String timestampInSeconds) {
        return unixTimeStampToLocalDateTime(Long.parseLong(timestampInSeconds));
    }

    public static long localDateTimeToUnixTimestamp(LocalDateTime time) {
        return time.atZone(TIMEZONE_ID).toEpochSecond();
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

    public static ZonedDateTime unixTimeStampToZonedDateTime(long timestampInSeconds) {
        return ZonedDateTime.ofInstant(Instant.ofEpochSecond(timestampInSeconds), TIMEZONE_ID);
    }
}

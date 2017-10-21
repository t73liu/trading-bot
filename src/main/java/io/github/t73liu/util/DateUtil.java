package io.github.t73liu.util;

import java.time.*;
import java.time.format.DateTimeFormatter;

public class DateUtil {
    public static final DateTimeFormatter LOCALDATE_ISO_FORMATTER = DateTimeFormatter.ISO_LOCAL_DATE;
    public static final DateTimeFormatter LOCALDATE_SHORT_FORMATTER = DateTimeFormatter.ofPattern("uuuuMMdd");
    public static final DateTimeFormatter LOCALDATETIME_ISO_FORMATTER = DateTimeFormatter.ISO_LOCAL_DATE_TIME;
    public static final String TIMEZONE = "America/New_York";
    public static final ZoneId TIMEZONE_ID = ZoneId.of(TIMEZONE);
    private static final int SECOND_TO_MILLISECOND = 1000;

    public static LocalDate getCurrentLocalDate() {
        return LocalDate.now(ZoneId.of(TIMEZONE));
    }

    public static LocalDateTime getCurrentLocalDateTime() {
        return LocalDateTime.now(ZoneId.of(TIMEZONE));
    }

    public static LocalDateTime unixSecondsToLocalDateTime(long timestampInSeconds) {
        return Instant.ofEpochSecond(timestampInSeconds)
                .atZone(TIMEZONE_ID)
                .toLocalDateTime();
    }

    public static LocalDateTime unixSecondsToLocalDateTime(String timestampInSeconds) {
        return unixSecondsToLocalDateTime(Long.parseLong(timestampInSeconds));
    }

    public static long localDateTimeToUnixSeconds(LocalDateTime time) {
        return localDateTimeToUnixMilliseconds(time) / SECOND_TO_MILLISECOND;
    }

    public static long localDateTimeToUnixMilliseconds(LocalDateTime time) {
        return time.atZone(TIMEZONE_ID).toInstant().toEpochMilli();
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

    public static ZonedDateTime unixSecondsToZonedDateTime(long timestampInSeconds) {
        return unixMillisecondsToZonedDateTime(timestampInSeconds * SECOND_TO_MILLISECOND);
    }

    public static ZonedDateTime unixMillisecondsToZonedDateTime(long timestampInMilliseconds) {
        return ZonedDateTime.ofInstant(Instant.ofEpochMilli(timestampInMilliseconds), TIMEZONE_ID);
    }
}

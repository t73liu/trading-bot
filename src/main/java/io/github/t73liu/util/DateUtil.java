package io.github.t73liu.util;

import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.time.format.DateTimeFormatter;

public class DateUtil {
    public static final DateTimeFormatter LOCALDATE_ISO_FORMATTER = DateTimeFormatter.ISO_LOCAL_DATE;
    public static final DateTimeFormatter LOCALDATE_SHORT_FORMATTER = DateTimeFormatter.ofPattern("uuuuMMdd");
    public static final DateTimeFormatter LOCALDATETIME_ISO_FORMATTER = DateTimeFormatter.ISO_LOCAL_DATE_TIME;
    public static final String TIMEZONE = "America/New_York";

    public static LocalDate getCurrentLocalDate() {
        return LocalDate.now(ZoneId.of(TIMEZONE));
    }

    public static LocalDateTime getCurrentLocalDateTime() {
        return LocalDateTime.now(ZoneId.of(TIMEZONE));
    }

    public static LocalDateTime convertFromUnixTimestamp(long timestampInSeconds) {
        return Instant.ofEpochSecond(timestampInSeconds)
                .atZone(ZoneId.of(TIMEZONE))
                .toLocalDateTime();
    }

    public static LocalDateTime convertFromUnixTimestamp(String timestampInSeconds) {
        return convertFromUnixTimestamp(Long.parseLong(timestampInSeconds));
    }

    public static long convertToUnixTimestamp(LocalDateTime time) {
        return time.atZone(ZoneId.of(TIMEZONE)).toEpochSecond();
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

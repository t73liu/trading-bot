package io.github.t73liu.util;

import org.junit.jupiter.api.Test;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.ZonedDateTime;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class DateUtilTest {
    @Test
    public void testLongUnixTimestampToLocalDate() {
        assertEquals(LocalDateTime.of(2017, 6, 28, 23, 22, 19), DateUtil.unixSecondsToLocalDateTime(1498706539));
    }

    @Test
    public void testStringUnixTimestampToLocalDate() {
        assertEquals(LocalDateTime.of(2017, 6, 28, 23, 22, 19), DateUtil.unixSecondsToLocalDateTime("1498706539"));
    }

    @Test
    public void testLocalDateTimeToUnixTimestamp() {
        assertEquals(1504577942, DateUtil.localDateTimeToUnixSeconds(LocalDateTime.of(2017, 9, 4, 22, 19, 2)));
    }

    @Test
    public void testISODateStringToLocalDate() {
        assertEquals(LocalDate.of(2017, 6, 30), DateUtil.parseLocalDateISO("2017-06-30"));
    }

    @Test
    public void testLocalDateToISODateString() {
        assertEquals("2017-06-30", DateUtil.formatLocalDateISO(LocalDate.of(2017, 6, 30)));
    }

    @Test
    public void testShortDateStringToLocalDate() {
        assertEquals(LocalDate.of(2017, 6, 30), DateUtil.parseLocalDateShort("20170630"));
    }

    @Test
    public void testLocalDateToShortDateString() {
        assertEquals("20170630", DateUtil.formatLocalDateShort(LocalDate.of(2017, 6, 30)));
    }

    @Test
    public void testUnixTimestampToZonedDateTime() {
        assertEquals(ZonedDateTime.of(2017, 10, 14, 12, 39, 36, 0, DateUtil.TIMEZONE_ID),
                DateUtil.unixSecondsToZonedDateTime(1507999176));
    }
}
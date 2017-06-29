package io.github.t73liu.util;

import org.junit.Test;

import java.time.LocalDate;
import java.time.LocalDateTime;

import static org.junit.Assert.assertEquals;

public class DateUtilTest {
    @Test
    public void testLongUnixTimestampToLocalDate() {
        assertEquals(LocalDateTime.of(2017, 6, 28, 23, 22, 19), DateUtil.convertUnixTimestamp(1498706539));
    }

    @Test
    public void testStringUnixTimestampToLocalDate() {
        assertEquals(LocalDateTime.of(2017, 6, 28, 23, 22, 19), DateUtil.convertUnixTimestamp("1498706539"));
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
}
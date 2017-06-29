package io.github.t73liu.util;

import org.junit.Test;

import java.time.LocalDate;

import static org.junit.Assert.*;

public class DateUtilTest {
    @Test
    public void testISODateStringToLocalDate() {
        assertEquals(LocalDate.of(2017, 6, 30), DateUtil.parseLocalDateISO("2017-06-30"));
    }

    @Test
    public void testLocalDateToISODateString() {
        assertEquals("2017-06-30", DateUtil.formatLocalDateISO(LocalDate.of(2017,6,30)));
    }

    @Test
    public void testShortDateStringToLocalDate() {
        assertEquals(LocalDate.of(2017, 6, 30), DateUtil.parseLocalDateShort("20170630"));
    }

    @Test
    public void testLocalDateToShortDateString() {
        assertEquals("20170630", DateUtil.formatLocalDateShort(LocalDate.of(2017,6,30)));
    }
}
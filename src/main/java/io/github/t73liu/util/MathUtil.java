package io.github.t73liu.util;

import eu.verdelhan.ta4j.Decimal;
import org.apache.commons.math3.util.Precision;

public class MathUtil {
    public static double roundDecimalToDouble(Decimal decimal) {
        return roundDecimalToDouble(decimal, 8);
    }

    public static double roundDecimalToDouble(Decimal decimal, int precision) {
        return Precision.round(decimal.toDouble(), precision);
    }
}

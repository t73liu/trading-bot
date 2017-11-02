package io.github.t73liu.util;

import org.apache.commons.math3.util.Precision;
import org.ta4j.core.Decimal;

public class MathUtil {
    public static double roundDecimalToDouble(Decimal decimal) {
        return roundDecimalToDouble(decimal, 8);
    }

    public static double roundDecimalToDouble(Decimal decimal, int precision) {
        return Precision.round(decimal.toDouble(), precision);
    }
}

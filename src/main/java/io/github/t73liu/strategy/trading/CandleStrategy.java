package io.github.t73liu.strategy.trading;

import eu.verdelhan.ta4j.*;
import eu.verdelhan.ta4j.indicators.candles.BearishEngulfingIndicator;
import eu.verdelhan.ta4j.indicators.candles.ThreeBlackCrowsIndicator;
import eu.verdelhan.ta4j.indicators.candles.ThreeWhiteSoldiersIndicator;
import eu.verdelhan.ta4j.trading.rules.BooleanIndicatorRule;

public class CandleStrategy {
    public static Strategy getStrategy(TimeSeries series) {
        Rule entryRule = new BooleanIndicatorRule(new ThreeWhiteSoldiersIndicator(series, 3, Decimal.valueOf(1.5)));
        Rule exitRule = new BooleanIndicatorRule(new BearishEngulfingIndicator(series))
                .or(new BooleanIndicatorRule(new ThreeBlackCrowsIndicator(series, 3, Decimal.valueOf(1.5))));
        return new BaseStrategy(entryRule, exitRule);
    }
}

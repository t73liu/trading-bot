package io.github.t73liu.strategy.trading;

import org.ta4j.core.*;
import org.ta4j.core.indicators.candles.BearishEngulfingIndicator;
import org.ta4j.core.indicators.candles.ThreeBlackCrowsIndicator;
import org.ta4j.core.indicators.candles.ThreeWhiteSoldiersIndicator;
import org.ta4j.core.indicators.helpers.ClosePriceIndicator;
import org.ta4j.core.trading.rules.BooleanIndicatorRule;
import org.ta4j.core.trading.rules.StopLossRule;

public class CandleStrategy {
    public static Strategy getStrategy(TimeSeries series) {
        Rule entryRule = new BooleanIndicatorRule(new ThreeWhiteSoldiersIndicator(series, 3, Decimal.valueOf(1.5)));
        Rule exitRule = new StopLossRule(new ClosePriceIndicator(series), Decimal.valueOf(5))
                .or(new BooleanIndicatorRule(new BearishEngulfingIndicator(series)))
                .or(new BooleanIndicatorRule(new ThreeBlackCrowsIndicator(series, 3, Decimal.valueOf(1.5))));
        return new BaseStrategy(entryRule, exitRule);
    }
}

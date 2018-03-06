package io.github.t73liu.strategy.trend;

import org.ta4j.core.BaseStrategy;
import org.ta4j.core.Rule;
import org.ta4j.core.Strategy;
import org.ta4j.core.TimeSeries;
import org.ta4j.core.indicators.CCIIndicator;
import org.ta4j.core.num.DoubleNum;
import org.ta4j.core.num.Num;
import org.ta4j.core.trading.rules.OverIndicatorRule;
import org.ta4j.core.trading.rules.UnderIndicatorRule;

public class CCIStrategy {
    public static Strategy getStrategy(TimeSeries series) {
        CCIIndicator longCci = new CCIIndicator(series, 200);
        CCIIndicator shortCci = new CCIIndicator(series, 5);
        Num plus100 = DoubleNum.valueOf(100);
        Num minus100 = DoubleNum.valueOf(-100);

        Rule entryRule = new OverIndicatorRule(longCci, plus100) // Bull trend
                .and(new UnderIndicatorRule(shortCci, minus100)); // Signal

        Rule exitRule = new UnderIndicatorRule(longCci, minus100) // Bear trend
                .and(new OverIndicatorRule(shortCci, plus100)); // Signal

        Strategy strategy = new BaseStrategy(entryRule, exitRule);
        strategy.setUnstablePeriod(5);
        return strategy;
    }
}

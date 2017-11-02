package io.github.t73liu.calc;

import org.springframework.stereotype.Service;
import org.ta4j.core.*;
import org.ta4j.core.analysis.criteria.TotalProfitCriterion;

@Service
public class EarningsCalculator {
    private static final AnalysisCriterion PROFIT_CRITERION = new TotalProfitCriterion();

    public static double calculateProfit(TimeSeries series, Strategy strategy, double takerFee) {
        TradingRecord tradingRecord = new TimeSeriesManager(series).run(strategy);
        return PROFIT_CRITERION.calculate(series, tradingRecord) - tradingRecord.getTradeCount() * takerFee;
    }
}

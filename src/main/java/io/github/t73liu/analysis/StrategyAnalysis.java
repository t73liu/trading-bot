package io.github.t73liu.analysis;

import it.unimi.dsi.fastutil.objects.Object2DoubleLinkedOpenHashMap;
import org.ta4j.core.*;
import org.ta4j.core.analysis.criteria.*;

import java.util.Map;

public class StrategyAnalysis {
    // TODO place into enum?
    private static final AnalysisCriterion AVERAGE_PROFIT_CRITERION = new AverageProfitCriterion();
    private static final AnalysisCriterion AVERAGE_PROFITABLE_TRADES_CRITERION = new AverageProfitableTradesCriterion();
    private static final AnalysisCriterion BUY_HOLD_CRITERION = new BuyAndHoldCriterion();
    private static final AnalysisCriterion MAX_DRAWDOWN_CRITERION = new MaximumDrawdownCriterion();
    private static final AnalysisCriterion NUMBER_BARS_CRITERION = new NumberOfBarsCriterion();
    private static final AnalysisCriterion NUMBER_TRADES_CRITERION = new NumberOfTradesCriterion();
    private static final AnalysisCriterion REWARD_RISK_RATIO_CRITERION = new RewardRiskRatioCriterion();
    private static final AnalysisCriterion PROFIT_CRITERION = new TotalProfitCriterion();
    private static final int TOTAL_CRITERION_COUNT = 9;
    private static final int INITIAL_AMOUNT = 1;

    public static Map<String, Double> analyze(TimeSeries series, Strategy strategy, double relativeTransactionFee, double flatTransactionFee) {
        TradingRecord tradingRecord = new TimeSeriesManager(series).run(strategy);
        Map<String, Double> analysisMap = new Object2DoubleLinkedOpenHashMap<>(TOTAL_CRITERION_COUNT);
        AnalysisCriterion transactionCostCriterion = new LinearTransactionCostCriterion(INITIAL_AMOUNT, relativeTransactionFee, flatTransactionFee);
        analysisMap.put("AVERAGE_PROFIT_CRITERION", AVERAGE_PROFIT_CRITERION.calculate(series, tradingRecord));
        analysisMap.put("AVERAGE_PROFITABLE_TRADES_CRITERION", AVERAGE_PROFITABLE_TRADES_CRITERION.calculate(series, tradingRecord));
        analysisMap.put("BUY_HOLD_CRITERION", BUY_HOLD_CRITERION.calculate(series, tradingRecord));
        analysisMap.put("MAX_DRAWDOWN_CRITERION", MAX_DRAWDOWN_CRITERION.calculate(series, tradingRecord));
        analysisMap.put("NUMBER_BARS_CRITERION", NUMBER_BARS_CRITERION.calculate(series, tradingRecord));
        analysisMap.put("NUMBER_TRADES_CRITERION", NUMBER_TRADES_CRITERION.calculate(series, tradingRecord));
        analysisMap.put("REWARD_RISK_RATIO_CRITERION", REWARD_RISK_RATIO_CRITERION.calculate(series, tradingRecord));
        analysisMap.put("TRANSACTION_COST_CRITERION", transactionCostCriterion.calculate(series, tradingRecord));
        analysisMap.put("PROFIT_CRITERION", PROFIT_CRITERION.calculate(series, tradingRecord) - analysisMap.getOrDefault("TRANSACTION_COST_CRITERION", 0d));
        return analysisMap;
    }
}

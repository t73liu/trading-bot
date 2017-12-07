package io.github.t73liu.analysis;

import io.github.t73liu.model.poloniex.PoloniexCandle;
import io.github.t73liu.strategy.trading.PlaceholderStrategy;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.ta4j.core.Bar;
import org.ta4j.core.BaseTimeSeries;
import org.ta4j.core.Strategy;
import org.ta4j.core.TimeSeries;

import java.util.List;
import java.util.Map;
import java.util.function.Function;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import static io.github.t73liu.util.MapperUtil.readCSV;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class StrategyAnalysisTest {
    private static final Logger LOGGER = LoggerFactory.getLogger(StrategyAnalysisTest.class);

    private static Stream<String> pathProvider() {
        return Stream.of("src/test/resources/USDT_ETH.300.poloniex.csv", "src/test/resources/USDT_ETH.900.poloniex.csv",
                "src/test/resources/USDT_ETH.1800.poloniex.csv", "src/test/resources/USDT_ETH.7200.poloniex.csv",
                "src/test/resources/USDT_XRP.300.poloniex.csv", "src/test/resources/USDT_XRP.900.poloniex.csv",
                "src/test/resources/USDT_XRP.1800.poloniex.csv", "src/test/resources/USDT_XRP.7200.poloniex.csv");
    }

    private static void analyzeStrategy(String path, Function<TimeSeries, Strategy> strategyFunction, String strategyName) throws Exception {
        List<Bar> ticks = readCSV(PoloniexCandle.class, path).stream()
                .map(PoloniexCandle::toTick)
                .collect(Collectors.toCollection(ObjectArrayList::new));
        TimeSeries series = new BaseTimeSeries(ticks);
        double relativeFee = 0.002;
        double flatFee = 0d;
        Map<String, Double> analysis = StrategyAnalysis.analyze(series, strategyFunction.apply(series), relativeFee, flatFee);
        LOGGER.info("Strategy: {}, Data: {}, Analysis: {}", strategyName, path, analysis);
        assertTrue(analysis.get("PROFIT_CRITERION") >= analysis.get("BUY_HOLD_CRITERION"));
    }

    @ParameterizedTest
    @MethodSource("pathProvider")
    public void testGlobalExtremeStrategy(String path) throws Exception {
        analyzeStrategy(path, PlaceholderStrategy::getGlobalExtremaStrategy, "GLOBAL EXTREMA");
    }

    @ParameterizedTest
    @MethodSource("pathProvider")
    public void testRsiStrategy(String path) throws Exception {
        analyzeStrategy(path, PlaceholderStrategy::getRsiStrategy, "RSI");
    }

    @ParameterizedTest
    @MethodSource("pathProvider")
    public void testCciStrategy(String path) throws Exception {
        analyzeStrategy(path, PlaceholderStrategy::getCciStrategy, "CCI");
    }

    @ParameterizedTest
    @MethodSource("pathProvider")
    public void testMovingMomentumStrategy(String path) throws Exception {
        analyzeStrategy(path, PlaceholderStrategy::getMovingMomentumStrategy, "MOVING MOMENTUM");
    }
}
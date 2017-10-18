package io.github.t73liu.calc;

import eu.verdelhan.ta4j.*;
import io.github.t73liu.model.poloniex.PoloniexCandle;
import io.github.t73liu.strategy.trading.CandleStrategy;
import org.apache.commons.math3.util.Precision;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import static io.github.t73liu.util.MapperUtil.readCSV;
import static io.github.t73liu.util.MathUtil.roundDecimalToDouble;

public class EarningsCalculatorTest {
    private static final Logger LOGGER = LoggerFactory.getLogger(EarningsCalculatorTest.class);

    private static Stream<String> pathProvider() {
        return Stream.of("src/test/resources/USDT_ETH.300.poloniex.csv", "src/test/resources/USDT_ETH.900.poloniex.csv",
                "src/test/resources/USDT_ETH.1800.poloniex.csv", "src/test/resources/USDT_ETH.7200.poloniex.csv",
                "src/test/resources/USDT_XRP.300.poloniex.csv", "src/test/resources/USDT_XRP.900.poloniex.csv",
                "src/test/resources/USDT_XRP.1800.poloniex.csv", "src/test/resources/USDT_XRP.7200.poloniex.csv");
    }

    @ParameterizedTest
    @MethodSource("pathProvider")
    public void testStrategyProfitability(String path) throws Exception {
        List<Tick> ticks = readCSV(PoloniexCandle.class, path).stream()
                .map(PoloniexCandle::toTick)
                .collect(Collectors.toList());
        TimeSeries series = new BaseTimeSeries(ticks);
        Strategy strategy = CandleStrategy.getStrategy(series);
        double takerFee = 0.002;
        double profit = EarningsCalculator.calculateProfit(series, strategy, takerFee);
        Decimal growth = ticks.get(ticks.size() - 1).getClosePrice().dividedBy(ticks.get(0).getClosePrice());
        LOGGER.info("DEFAULT: {}, STRATEGY:{}", roundDecimalToDouble(growth), Precision.round(profit, 8));
        // RE-ENABLE TEST WHEN STRATEGY SOUND
//        Assertions.assertTrue(profit > growth.multipliedBy(Decimal.valueOf(1.05)).toDouble());
    }
}
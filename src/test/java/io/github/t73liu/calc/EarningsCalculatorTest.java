package io.github.t73liu.calc;

import eu.verdelhan.ta4j.BaseTimeSeries;
import eu.verdelhan.ta4j.Strategy;
import eu.verdelhan.ta4j.Tick;
import eu.verdelhan.ta4j.TimeSeries;
import io.github.t73liu.exchange.poloniex.rest.PoloniexMarketService;
import io.github.t73liu.model.poloniex.PoloniexCandle;
import io.github.t73liu.strategy.trading.CandleStrategy;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.MethodSource;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import static io.github.t73liu.util.MapperUtil.readCSV;

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
                .map(PoloniexMarketService::mapExchangeCandleToTick)
                .collect(Collectors.toList());
        TimeSeries series = new BaseTimeSeries(ticks);
        Strategy strategy = CandleStrategy.getStrategy(series);
        double takerFee = 0.002;
        double profit = EarningsCalculator.calculateProfit(series, strategy, takerFee);
        if (profit > 1.05) {
            LOGGER.info("File:{}, PROFIT:{}", path, profit);
        } else {
            LOGGER.error("File:{}, LOSS:{}", path, profit);
        }
//        Assertions.assertTrue(profit > 1.05);
    }
}
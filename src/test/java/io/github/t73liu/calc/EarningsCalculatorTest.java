package io.github.t73liu.calc;

import eu.verdelhan.ta4j.BaseTimeSeries;
import eu.verdelhan.ta4j.Strategy;
import eu.verdelhan.ta4j.Tick;
import eu.verdelhan.ta4j.TimeSeries;
import io.github.t73liu.exchange.poloniex.PoloniexService;
import io.github.t73liu.model.poloniex.PoloniexCandle;
import io.github.t73liu.strategy.trading.CandleStrategy;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;
import java.util.zip.GZIPInputStream;

import static io.github.t73liu.util.ObjectMapperFactory.OBJECT_READER;

public class EarningsCalculatorTest {
    private static final Logger LOGGER = LoggerFactory.getLogger(EarningsCalculatorTest.class);

    @Test
    public void testStrategyProfitability() throws Exception {
        PoloniexCandle[] data = OBJECT_READER.forType(PoloniexCandle[].class).readValue(new GZIPInputStream(Files.newInputStream(Paths.get("src/test/resources/USDT_BTC.poloniex.json.gz"))));
        List<Tick> ticks = Arrays.stream(data).map(PoloniexService::mapExchangeCandleToTick).collect(Collectors.toList());
        TimeSeries series = new BaseTimeSeries(ticks);
        Strategy strategy = CandleStrategy.getStrategy(series);
        double takerFee = 0.002;
        double profit = EarningsCalculator.calculateProfit(series, strategy, takerFee);
        LOGGER.info("Profit:{}", profit);
        Assertions.assertTrue(profit > 1.1);
    }
}
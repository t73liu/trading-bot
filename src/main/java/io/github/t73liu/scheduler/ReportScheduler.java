package io.github.t73liu.scheduler;

import eu.verdelhan.ta4j.Tick;
import io.github.t73liu.exchange.bittrex.rest.BittrexService;
import io.github.t73liu.exchange.poloniex.rest.PoloniexService;
import io.github.t73liu.exchange.quadriga.rest.QuadrigaService;
import io.github.t73liu.model.Balance;
import io.github.t73liu.model.CandlestickIntervals;
import io.github.t73liu.report.MailingService;
import io.github.t73liu.util.DateUtil;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.stream.Collectors;

import static io.github.t73liu.model.poloniex.PoloniexPair.USDT_XRP;
import static io.github.t73liu.util.DateUtil.getCurrentLocalDateTime;

@Component
public class ReportScheduler {
    private static final Logger LOGGER = LoggerFactory.getLogger(ReportScheduler.class);

    private final BittrexService bittrexService;

    private final PoloniexService poloniexService;

    private final QuadrigaService quadrigaService;

    private final MailingService mailingService;

    @Autowired
    public ReportScheduler(BittrexService bittrexService, PoloniexService poloniexService,
                           QuadrigaService quadrigaService, MailingService mailingService) {
        this.bittrexService = bittrexService;
        this.poloniexService = poloniexService;
        this.quadrigaService = quadrigaService;
        this.mailingService = mailingService;
    }

    @Scheduled(cron = "${schedules.report.cron:0 0 16 * * *}", zone = DateUtil.TIMEZONE)
    public void createDailyReport() {
        // TODO implement daily report generation
        LOGGER.info("Creating Daily Report");
    }

    private BigDecimal lastBuyRate = null;

    //    @Scheduled(fixedDelay = 5000, zone = DateUtil.TIMEZONE)
    public void checkCandlesticks() throws Exception {
        LOGGER.info("Checking Poloniex Candlesticks for opportunities");
        Map<String, Map<String, String>> allBalances = poloniexService.getCompleteBalances();
        BigDecimal availableUSDT = new BigDecimal(allBalances.get("USDT").get("available")).setScale(8, RoundingMode.HALF_UP);
        BigDecimal availableXRP = new BigDecimal(allBalances.get("XRP").get("available")).setScale(8, RoundingMode.HALF_UP);

        BigDecimal sellRate = new BigDecimal(poloniexService.getTickers().get("USDT_XRP").get("highestBid")).setScale(8, RoundingMode.HALF_UP);
        BigDecimal buyRate = new BigDecimal(poloniexService.getTickers().get("USDT_XRP").get("lowestAsk")).setScale(8, RoundingMode.HALF_UP);

        LocalDateTime endLocalDateTime = getCurrentLocalDateTime();
        LocalDateTime startLocalDateTime = endLocalDateTime.minusHours(1);
        List<Tick> sticks = poloniexService.getCandlestick(USDT_XRP, startLocalDateTime, endLocalDateTime, CandlestickIntervals.FIFTEEN_MIN);
        LOGGER.info("Resulting candlesticks: {}", sticks);
    }

    //    @Scheduled(fixedDelay = 7200000, zone = DateUtil.TIMEZONE)
    public void reportBalances() throws Exception {
        LOGGER.info("Reporting Poloniex Balance values");
        Map<String, Map<String, String>> allBalances = poloniexService.getCompleteBalances();
        double usdtRate = Double.valueOf(poloniexService.getTickers().get("USDT_BTC").get("last"));
        List<Balance> balanceList = new ObjectArrayList<>(2);
        Balance xrp = new Balance();
        xrp.setAvailable(new BigDecimal(allBalances.get("XRP").get("available")).setScale(8, RoundingMode.HALF_UP));
        xrp.setOnOrders(new BigDecimal(allBalances.get("XRP").get("onOrders")).setScale(8, RoundingMode.HALF_UP));
        xrp.setCurrency("XRP");
        xrp.setUsdValue(new BigDecimal(Double.valueOf(allBalances.get("XRP").get("btcValue")) * usdtRate));
        balanceList.add(xrp);
        Balance usdt = new Balance();
        usdt.setAvailable(new BigDecimal(allBalances.get("USDT").get("available")).setScale(8, RoundingMode.HALF_UP));
        usdt.setOnOrders(new BigDecimal(allBalances.get("USDT").get("onOrders")).setScale(8, RoundingMode.HALF_UP));
        usdt.setCurrency("USDT");
        usdt.setUsdValue(usdt.getAvailable());
        balanceList.add(usdt);
        Balance btcd = new Balance();
        btcd.setAvailable(new BigDecimal(allBalances.get("BTCD").get("available")).setScale(8, RoundingMode.HALF_UP));
        btcd.setOnOrders(new BigDecimal(allBalances.get("BTCD").get("onOrders")).setScale(8, RoundingMode.HALF_UP));
        btcd.setCurrency("BTCD");
        xrp.setUsdValue(new BigDecimal(Double.valueOf(allBalances.get("BTCD").get("btcValue")) * usdtRate));
        balanceList.add(btcd);
        Optional<BigDecimal> total = balanceList.stream().map(Balance::getUsdValue).collect(Collectors.reducing(BigDecimal::add));
        mailingService.sendMail("Poloniex Balance", "Total: " + total.get() + " " + String.valueOf(balanceList));
    }
}

package io.github.t73liu.schedules;

import io.github.t73liu.model.Balance;
import io.github.t73liu.service.BittrexService;
import io.github.t73liu.service.MailingService;
import io.github.t73liu.service.PoloniexService;
import io.github.t73liu.service.QuadrigaService;
import io.github.t73liu.util.DateUtil;
import it.unimi.dsi.fastutil.objects.ObjectArrayList;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import java.math.BigDecimal;
import java.math.RoundingMode;
import java.util.List;
import java.util.Map;

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

    @Scheduled(fixedDelay = 72000, zone = DateUtil.TIMEZONE)
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
        allBalances.get("USDT");
        Balance usdt = new Balance();
        usdt.setAvailable(new BigDecimal(allBalances.get("USDT").get("available")).setScale(8, RoundingMode.HALF_UP));
        usdt.setOnOrders(new BigDecimal(allBalances.get("USDT").get("onOrders")).setScale(8, RoundingMode.HALF_UP));
        usdt.setCurrency("USDT");
        usdt.setUsdValue(usdt.getAvailable().add(usdt.getOnOrders()));
        balanceList.add(usdt);
        mailingService.sendMail("Poloniex Balance", String.valueOf(balanceList));
    }
}

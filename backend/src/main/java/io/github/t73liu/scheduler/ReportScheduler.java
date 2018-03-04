package io.github.t73liu.scheduler;

import io.github.t73liu.exchange.bittrex.rest.BittrexAccountService;
import io.github.t73liu.exchange.poloniex.rest.PoloniexAccountService;
import io.github.t73liu.exchange.quadriga.rest.QuadrigaAccountService;
import io.github.t73liu.report.MailingService;
import io.github.t73liu.util.DateUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

@Component
public class ReportScheduler {
    private static final Logger LOGGER = LoggerFactory.getLogger(ReportScheduler.class);

    private final BittrexAccountService bittrexAccountService;

    private final PoloniexAccountService poloniexAccountService;

    private final QuadrigaAccountService quadrigaAccountService;

    private final MailingService mailingService;

    @Autowired
    public ReportScheduler(BittrexAccountService bittrexAccountService, PoloniexAccountService poloniexAccountService,
                           QuadrigaAccountService quadrigaAccountService, MailingService mailingService) {
        this.bittrexAccountService = bittrexAccountService;
        this.poloniexAccountService = poloniexAccountService;
        this.quadrigaAccountService = quadrigaAccountService;
        this.mailingService = mailingService;
    }

    @Scheduled(cron = "${schedules.report.cron:0 0 16 * * *}", zone = DateUtil.TIMEZONE)
    public void createDailyReport() {
        // TODO implement daily report generation
        LOGGER.info("Creating Daily Report");
    }
}

package io.github.t73liu.schedules;

import io.github.t73liu.util.DateUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

@Component
public class ReportScheduler {
    private final Logger LOGGER = LoggerFactory.getLogger(this.getClass());

    @Scheduled(cron = "${schedules.report.cron:0 0 16 * * *}", zone = DateUtil.TIMEZONE)
    public void createDailyReport() {
        // TODO implement daily report generation
        LOGGER.info("Creating Daily Report");
    }
}

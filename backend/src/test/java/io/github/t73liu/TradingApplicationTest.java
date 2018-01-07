package io.github.t73liu;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.junit.jupiter.SpringJUnitConfig;

@SpringBootTest
@SpringJUnitConfig
@ActiveProfiles("test")
class TradingApplicationTest {
    @Test
    void testStartup() {
    }
}

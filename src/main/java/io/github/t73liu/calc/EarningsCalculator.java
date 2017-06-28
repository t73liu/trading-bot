package io.github.t73liu.calc;

import io.github.t73liu.model.Order;
import org.springframework.stereotype.Component;

@Component
public class EarningsCalculator {
    public Double calculateProfit(Order buy, Order sale) {
        return 0d;
    }
}

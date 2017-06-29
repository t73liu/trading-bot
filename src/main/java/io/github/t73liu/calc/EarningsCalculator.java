package io.github.t73liu.calc;

import io.github.t73liu.model.Order;
import org.springframework.stereotype.Component;

@Component
public class EarningsCalculator {
    // FIXME fix logic
    public Double calculateProfit(Order buy, Order sale) {
        return Math.min(buy.getQuantity(), sale.getQuantity()) * (sale.getPrice() - buy.getPrice());
    }
}

package io.github.t73liu.calc;

import io.github.t73liu.model.Order;
import org.apache.commons.math3.util.FastMath;
import org.springframework.stereotype.Service;

@Service
public class EarningsCalculator {
    public Double calculateProfit(Order buy, Order sale) {
        return FastMath.min(buy.getQuantity(), sale.getQuantity()) * (sale.getPrice() - buy.getPrice());
    }
}

package io.github.t73liu.service;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Service;

@Service
@ConfigurationProperties(prefix = "quadrigacx")
public class QuadrigaService extends ExchangeService {
}

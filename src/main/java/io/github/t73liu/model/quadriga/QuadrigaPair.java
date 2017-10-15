package io.github.t73liu.model.quadriga;

public enum QuadrigaPair {
    BCH_CAD("bch_cad"),
    BTC_CAD("btc_cad"),
    BTC_USD("btc_usd"),
    ETH_BTC("eth_btc"),
    ETH_CAD("eth_cad"),
    LTC_CAD("ltc_cad");

    private final String pairName;

    QuadrigaPair(String pairName) {
        this.pairName = pairName;
    }

    public String getPairName() {
        return pairName;
    }
}

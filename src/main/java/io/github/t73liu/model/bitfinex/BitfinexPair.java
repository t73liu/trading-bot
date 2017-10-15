package io.github.t73liu.model.bitfinex;

public enum BitfinexPair {
    tBTCUSD("btcusd"),
    tLTCUSD("ltcusd"),
    tLTCBTC("ltcbtc"),
    tETHUSD("ethusd"),
    tETHBTC("ethbtc"),
    tETCBTC("etcbtc"),
    tETCUSD("etcusd"),
    tRRTUSD("rrtusd"),
    tRRTBTC("rrtbtc"),
    tZECUSD("zecusd"),
    tZECBTC("zecbtc"),
    tXMRUSD("xmrusd"),
    tXMRBTC("xmrbtc"),
    tDSHUSD("dshusd"),
    tDSHBTC("dshbtc"),
    tBCCBTC("bccbtc"),
    tBCUBTC("bcubtc"),
    tBCCUSD("bccusd"),
    tBCUUSD("bcuusd"),
    tXRPUSD("xrpusd"),
    tXRPBTC("xrpbtc"),
    tIOTUSD("iotusd"),
    tIOTBTC("iotbtc"),
    tIOTETH("ioteth"),
    tEOSUSD("eosusd"),
    tEOSBTC("eosbtc"),
    tEOSETH("eoseth"),
    tSANUSD("sanusd"),
    tSANBTC("sanbtc"),
    tSANETH("saneth"),
    tOMGUSD("omgusd"),
    tOMGBTC("omgbtc"),
    tOMGETH("omgeth"),
    tBCHUSD("bchusd"),
    tBCHBTC("bchbtc"),
    tBCHETH("bcheth"),
    tNEOUSD("neousd"),
    tNEOBTC("neobtc"),
    tNEOETH("neoeth"),
    tETPUSD("etpusd"),
    tETPBTC("etpbtc"),
    tETPETH("etpeth"),
    tQTMUSD("qtmusd"),
    tQTMBTC("qtmbtc"),
    tQTMETH("qtmeth"),
    tBT1USD("bt1usd"),
    tBT2USD("bt2usd"),
    tBT1BTC("bt1btc"),
    tBT2BTC("bt2btc"),
    tAVTUSD("avtusd"),
    tAVTBTC("avtbtc"),
    tAVTETH("avteth");

    private final String pairName;

    BitfinexPair(String pairName) {
        this.pairName = pairName;
    }

    public String getPairName() {
        return pairName;
    }
}

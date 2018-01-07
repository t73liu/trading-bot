package io.github.t73liu.model.bitfinex;

public enum BitfinexPair {
    BTC_USD("tBTCUSD"),
    LTC_USD("tLTCUSD"),
    LTC_BTC("tLTCBTC"),
    ETH_USD("tETHUSD"),
    ETH_BTC("tETHBTC"),
    ETC_BTC("tETCBTC"),
    ETC_USD("tETCUSD"),
    RRT_USD("tRRTUSD"),
    RRT_BTC("tRRTBTC"),
    ZEC_USD("tZECUSD"),
    ZEC_BTC("tZECBTC"),
    XMR_USD("tXMRUSD"),
    XMR_BTC("tXMRBTC"),
    DSH_USD("tDSHUSD"),
    DSH_BTC("tDSHBTC"),
    BCC_BTC("tBCCBTC"),
    BCU_BTC("tBCUBTC"),
    BCC_USD("tBCCUSD"),
    BCU_USD("tBCUUSD"),
    XRP_USD("tXRPUSD"),
    XRP_BTC("tXRPBTC"),
    IOT_USD("tIOTUSD"),
    IOT_BTC("tIOTBTC"),
    IOT_ETH("tIOTETH"),
    EOS_USD("tEOSUSD"),
    EOS_BTC("tEOSBTC"),
    EOS_ETH("tEOSETH"),
    SAN_USD("tSANUSD"),
    SAN_BTC("tSANBTC"),
    SAN_ETH("tSANETH"),
    OMG_USD("tOMGUSD"),
    OMG_BTC("tOMGBTC"),
    OMG_ETH("tOMGETH"),
    BCH_USD("tBCHUSD"),
    BCH_BTC("tBCHBTC"),
    BCH_ETH("tBCHETH"),
    NEO_USD("tNEOUSD"),
    NEO_BTC("tNEOBTC"),
    NEO_ETH("tNEOETH"),
    ETP_USD("tETPUSD"),
    ETP_BTC("tETPBTC"),
    ETP_ETH("tETPETH"),
    QTM_USD("tQTMUSD"),
    QTM_BTC("tQTMBTC"),
    QTM_ETH("tQTMETH"),
    BT1_USD("tBT1USD"),
    BT2_USD("tBT2USD"),
    BT1_BTC("tBT1BTC"),
    BT2_BTC("tBT2BTC"),
    AVT_USD("tAVTUSD"),
    AVT_BTC("tAVTBTC"),
    AVT_ETH("tAVTETH");

    private final String pairName;

    BitfinexPair(String pairName) {
        this.pairName = pairName;
    }

    public String getPairName() {
        return pairName;
    }
}

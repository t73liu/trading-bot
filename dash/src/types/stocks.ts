export interface Stock {
  symbol: string;
  company: string;
}

export interface StockInfo {
  symbol: string;
  currentVolume: number;
  lastCandlePrice: number;
  info: unknown;
}

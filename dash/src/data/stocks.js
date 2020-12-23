import { stringify } from "query-string";
import { getJSON } from "../utils/http";

export const fetchStocks = () => getJSON("/api/stocks");

export const fetchStockInfo = (symbol) => getJSON(`/api/stocks/${symbol}`);

export const fetchStockNews = (symbol) => getJSON(`/api/stocks/${symbol}/news`);

export const fetchStockCharts = (symbol, candleSize, showExtendedHours) => {
  const query = stringify({
    interval: "intraday",
    candleSize,
    showExtendedHours,
  });
  return getJSON(`/api/stocks/${symbol}/charts?${query}`);
};

export const fetchGapStocks = () => getJSON("/api/stocks/gaps");

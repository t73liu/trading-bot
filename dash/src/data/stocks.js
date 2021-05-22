import { stringify } from "query-string";
import { fetchJSON } from "../utils/http";

export const fetchStocks = () => fetchJSON("/api/stocks");

export const fetchStockInfo = (symbol) => fetchJSON(`/api/stocks/${symbol}`);

export const fetchStockCharts = (symbol, candleSize, showExtendedHours) => {
  const query = stringify({
    interval: "intraday",
    candleSize,
    showExtendedHours,
  });
  return fetchJSON(`/api/stocks/${symbol}/charts?${query}`);
};

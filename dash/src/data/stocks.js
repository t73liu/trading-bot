import { stringify } from "query-string";
import { getJSON } from "../utils/http";

export const fetchStockInfo = (symbol) => getJSON(`/api/stocks/${symbol}`);

export const fetchStockNews = (symbol) => getJSON(`/api/stocks/${symbol}/news`);

export const fetchStockCandles = (symbol, showExtendedHours) => {
  const query = stringify({
    interval: "intraday",
    showExtendedHours,
  });
  return getJSON(`/api/stocks/${symbol}/candles?${query}`);
};

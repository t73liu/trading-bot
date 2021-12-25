import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

import { Stock, StockInfo } from "../types/stocks";

// Define a service using a base URL and expected endpoints
const stockAPI = createApi({
  reducerPath: "stocks",
  baseQuery: fetchBaseQuery({ baseUrl: "/api" }),
  endpoints: (builder) => ({
    getStocks: builder.query<Stock[], void>({
      query: () => "/stocks",
    }),
    getStockInfo: builder.query<StockInfo, string>({
      query: (symbol: string) => `/stocks/${symbol}`,
    }),
  }),
});

// Export hooks for usage in function components, which are
// auto-generated based on the defined endpoints
export const { useGetStocksQuery, useGetStockInfoQuery } = stockAPI;

export default stockAPI;

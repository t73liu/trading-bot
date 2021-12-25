import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

import { Watchlist } from "../types/watchlists";

const watchlistAPI = createApi({
  reducerPath: "watchlists",
  baseQuery: fetchBaseQuery({ baseUrl: "/api" }),
  endpoints: (builder) => ({
    getWatchlists: builder.query<Watchlist[], void>({
      query: () => "/account/watchlists",
    }),
  }),
});

export const { useGetWatchlistsQuery } = watchlistAPI;

export default watchlistAPI;

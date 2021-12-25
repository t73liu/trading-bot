import { configureStore } from "@reduxjs/toolkit";
import { setupListeners } from "@reduxjs/toolkit/query/react";

import stockAPI from "./stocks";
import watchlistAPI from "./watchlists";

const store = configureStore({
  reducer: {
    [stockAPI.reducerPath]: stockAPI.reducer,
    [watchlistAPI.reducerPath]: watchlistAPI.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(stockAPI.middleware, watchlistAPI.middleware),
});

setupListeners(store.dispatch);

// Infer the `RootState` and `AppDispatch` types from the store itself.
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;

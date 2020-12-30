import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";
import logger from "redux-logger";
import settingsReducer from "./settings";
import stocksReducer from "./stocks";
import watchlistsReducer from "./watchlists";

const store = configureStore({
  reducer: {
    settings: settingsReducer,
    stocks: stocksReducer,
    watchlists: watchlistsReducer,
  },
  middleware: [...getDefaultMiddleware(), logger],
  devTools: process.env.NODE_ENV !== "production",
});

export default store;

import { configureStore, getDefaultMiddleware } from "@reduxjs/toolkit";
import logger from "redux-logger";
import accountReducer from "./account";
import stocksReducer from "./stocks";

const store = configureStore({
  reducer: {
    account: accountReducer,
    stocks: stocksReducer,
  },
  middleware: [...getDefaultMiddleware(), logger],
  devTools: process.env.NODE_ENV !== "production",
});

export default store;

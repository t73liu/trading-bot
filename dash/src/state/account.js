import { createSlice } from "@reduxjs/toolkit";

const accountSlice = createSlice({
  name: "account",
  initialState: {
    watchlists: {},
    showExtendedHours: false,
    candleSize: "1min",
  },
  reducers: {
    createWatchlist: (draft, { payload }) => {
      draft.watchlists[payload] = [];
    },
    deleteWatchlist: (draft, { payload }) => {
      delete draft.watchlists[payload];
    },
    addStockToWatchlist: (draft, { payload }) => {
      const { watchlistID, stock } = payload;
      const stocks = draft.watchlists[watchlistID];
      if (!stocks.includes((s) => s === stock)) {
        draft.watchlists[watchlistID].push(stock);
      }
    },
    removeStockFromWatchlist: (draft, { payload }) => {
      const { watchlistID, stock } = payload;
      const stocks = draft.watchlists[watchlistID];
      draft.watchlists[watchlistID] = stocks.filter((s) => s !== stock);
    },
    toggleShowExtendedHours: (draft) => {
      draft.showExtendedHours = !draft.showExtendedHours;
    },
    setCandleSize: (draft, { payload }) => {
      draft.candleSize = payload;
    },
  },
});

export const {
  createWatchlist,
  deleteWatchlist,
  addStockToWatchlist,
  removeStockFromWatchlist,
  toggleShowExtendedHours,
  setCandleSize,
} = accountSlice.actions;

export default accountSlice.reducer;

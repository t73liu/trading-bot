import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { fetchStocks } from "../data/stocks";

export const fetchStocksThunk = createAsyncThunk(
  "stocks/fetchStocks",
  async () => {
    const response = await fetchStocks();
    if (response instanceof Error) {
      throw response;
    }
    return response;
  }
);

const stocksSlice = createSlice({
  name: "stocks",
  initialState: {
    allStocks: [],
    isLoading: false,
  },
  reducers: {},
  extraReducers: {
    [fetchStocksThunk.pending]: (draft) => {
      draft.isLoading = true;
    },
    [fetchStocksThunk.fulfilled]: (draft, { payload }) => {
      draft.isLoading = false;
      draft.allStocks = payload;
    },
    [fetchStocksThunk.rejected]: (draft) => {
      draft.isLoading = false;
    },
  },
});

export const { createWatchlist, deleteWatchlist } = stocksSlice.actions;

export default stocksSlice.reducer;

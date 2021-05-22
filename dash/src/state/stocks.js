import {
  createAsyncThunk,
  createSlice,
  createEntityAdapter,
} from "@reduxjs/toolkit";
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

const stocksAdapter = createEntityAdapter();

const initialState = stocksAdapter.getInitialState({ isLoading: false });

const stocksSlice = createSlice({
  name: "stocks",
  initialState,
  reducers: {},
  extraReducers: {
    [fetchStocksThunk.pending]: (draft) => {
      draft.isLoading = true;
    },
    [fetchStocksThunk.fulfilled]: (draft, action) => {
      stocksAdapter.upsertMany(draft, action);
      draft.isLoading = false;
    },
    [fetchStocksThunk.rejected]: (draft) => {
      draft.isLoading = false;
    },
  },
});

export const { selectAll: selectAllStocks, selectById: selectStockByID } =
  stocksAdapter.getSelectors((state) => state.stocks);

export default stocksSlice.reducer;

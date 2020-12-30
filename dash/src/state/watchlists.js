import {
  createAsyncThunk,
  createEntityAdapter,
  createSlice,
} from "@reduxjs/toolkit";
import { fetchWatchlists } from "../data/watchlists";

export const fetchWatchlistsThunk = createAsyncThunk(
  "account/fetchWatchlists",
  async () => {
    const response = await fetchWatchlists();
    if (response instanceof Error) {
      throw response;
    }
    return response;
  }
);

const watchlistsAdapter = createEntityAdapter();

const initialState = watchlistsAdapter.getInitialState({ isLoading: false });

const watchlistsSlice = createSlice({
  name: "watchlists",
  initialState,
  reducers: {
    createWatchlist: watchlistsAdapter.addOne,
    deleteWatchlist: watchlistsAdapter.removeOne,
    updateWatchlist: watchlistsAdapter.updateOne,
  },
  extraReducers: {
    [fetchWatchlistsThunk.pending]: (draft) => {
      draft.isLoading = true;
    },
    [fetchWatchlistsThunk.fulfilled]: (draft, action) => {
      watchlistsAdapter.upsertMany(draft, action);
      draft.isLoading = false;
    },
    [fetchWatchlistsThunk.rejected]: (draft) => {
      draft.isLoading = false;
    },
  },
});

export const {
  selectAll: selectAllWatchlists,
} = watchlistsAdapter.getSelectors((state) => state.watchlists);

export default watchlistsSlice.reducer;

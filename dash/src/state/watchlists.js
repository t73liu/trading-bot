import {
  createAsyncThunk,
  createEntityAdapter,
  createSlice,
} from "@reduxjs/toolkit";
import {
  createWatchlist,
  deleteWatchlist,
  fetchWatchlists,
  updateWatchlist,
} from "../data/watchlists";

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

export const updateWatchlistThunk = createAsyncThunk(
  "account/updateWatchlistThunk",
  async (watchlist) => {
    const response = await updateWatchlist(watchlist);
    if (response instanceof Error) {
      throw response;
    }
    return response;
  }
);

export const deleteWatchlistThunk = createAsyncThunk(
  "account/deleteWatchlistThunk",
  async (id) => {
    const response = await deleteWatchlist(id);
    if (response instanceof Error) {
      throw response;
    }
    return id;
  }
);

export const createWatchlistThunk = createAsyncThunk(
  "account/createWatchlistThunk",
  async () => {
    const response = await createWatchlist();
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
    watchStock: (draft, { payload }) => {
      const { watchlistID, stockID } = payload;
      const existingStockIDs = draft.entities[watchlistID].stockIDs;
      const updatedStockIDs = new Set([...existingStockIDs, stockID]);
      watchlistsAdapter.updateOne(draft, {
        id: watchlistID,
        changes: {
          stockIDs: [...updatedStockIDs],
        },
      });
    },
    unwatchStock: (draft, { payload }) => {
      const { watchlistID, stockID } = payload;
      const existingStockIDs = draft.entities[watchlistID].stockIDs;
      watchlistsAdapter.updateOne(draft, {
        id: watchlistID,
        changes: {
          stockIDs: existingStockIDs.filter((id) => id !== stockID),
        },
      });
    },
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
    [updateWatchlistThunk().pending]: (draft) => {
      draft.isLoading = true;
    },
    [updateWatchlistThunk.fulfilled]: (draft) => {
      draft.isLoading = false;
    },
    [updateWatchlistThunk.rejected]: (draft) => {
      draft.isLoading = false;
    },
    [deleteWatchlistThunk.pending]: (draft) => {
      draft.isLoading = true;
    },
    [deleteWatchlistThunk.fulfilled]: (draft, action) => {
      watchlistsAdapter.removeOne(draft, action);
      draft.isLoading = false;
    },
    [deleteWatchlistThunk.rejected]: (draft) => {
      draft.isLoading = false;
    },
    [createWatchlistThunk.pending]: (draft) => {
      draft.isLoading = true;
    },
    [createWatchlistThunk.fulfilled]: (draft, action) => {
      watchlistsAdapter.upsertOne(draft, action);
      draft.isLoading = false;
    },
    [createWatchlistThunk.rejected]: (draft) => {
      draft.isLoading = false;
    },
  },
});

export const {
  selectAll: selectAllWatchlists,
} = watchlistsAdapter.getSelectors((state) => state.watchlists);

export const { watchStock, unwatchStock } = watchlistsSlice.actions;

export default watchlistsSlice.reducer;

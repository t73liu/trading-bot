import { createSlice } from "@reduxjs/toolkit";

const accountSlice = createSlice({
  name: "account",
  initialState: {
    watchlists: {},
  },
  reducers: {
    addStock: (draft, { payload }) => {
      draft.watchlists = [payload];
    },
  },
});

export default accountSlice.reducer;

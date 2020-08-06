import { createSlice } from "@reduxjs/toolkit";

const accountSlice = createSlice({
  name: "account",
  initialState: {
    watchlists: {},
    showExtendedHours: false,
  },
  reducers: {
    addStock: (draft, { payload }) => {
      draft.watchlists = [payload];
    },
    toggleShowExtendedHours: (draft) => {
      draft.showExtendedHours = !draft.showExtendedHours;
    },
  },
});

export default accountSlice.reducer;

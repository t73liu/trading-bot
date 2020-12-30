import { createSelector, createSlice } from "@reduxjs/toolkit";

export const getCandleSize = createSelector(
  (state) => state.settings,
  (account) => account.candleSize
);

export const getShowExtendedHours = createSelector(
  (state) => state.settings,
  (account) => account.showExtendedHours
);

const settingsSlice = createSlice({
  name: "settings",
  initialState: {
    showExtendedHours: false,
    candleSize: "1min",
  },
  reducers: {
    toggleShowExtendedHours: (draft) => {
      draft.showExtendedHours = !draft.showExtendedHours;
    },
    setCandleSize: (draft, { payload }) => {
      draft.candleSize = payload;
    },
  },
});

export const { toggleShowExtendedHours, setCandleSize } = settingsSlice.actions;

export default settingsSlice.reducer;

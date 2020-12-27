import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { createSelector } from "@reduxjs/toolkit";
import { fetchWatchlistsThunk } from "../../state/account";

const getWatchlists = createSelector(
  (state) => state.account,
  (account) => account.watchlists
);

const Watchlists = () => {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchWatchlistsThunk());
  }, [dispatch]);
  const watchlists = useSelector(getWatchlists);
  return (
    <div>
      <h2>Watchlists</h2>
      <ul>
        {watchlists.map((watchlist) => (
          <li key={watchlist.id}>{watchlist.name}</li>
        ))}
      </ul>
    </div>
  );
};

export default Watchlists;

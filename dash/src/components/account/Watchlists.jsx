import React, { useCallback, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Button } from "@material-ui/core";
import {
  createWatchlistThunk,
  fetchWatchlistsThunk,
  selectAllWatchlists,
} from "../../state/watchlists";
import { useTitleContext } from "../../state/title-context";
import Watchlist from "./Watchlist";

const Watchlists = () => {
  const { setTitle } = useTitleContext();
  useEffect(() => setTitle("Watchlists"), [setTitle]);
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchWatchlistsThunk());
  }, [dispatch]);
  const watchlists = useSelector(selectAllWatchlists);
  const handleCreateWatchlist = useCallback(() => {
    dispatch(createWatchlistThunk());
  }, [dispatch]);
  return (
    <div>
      {watchlists?.map(({ id, name, stockIDs }) => (
        <Watchlist key={id} id={id} name={name} stockIDs={stockIDs} />
      ))}
      <Button variant="contained" onClick={handleCreateWatchlist}>
        New
      </Button>
    </div>
  );
};

export default Watchlists;

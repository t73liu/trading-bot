import React, { useCallback, useEffect, useState } from "react";
import { Button } from "@material-ui/core";
import PropTypes from "prop-types";
import { updateWatchlist } from "../../data/watchlists";

const Watchlist = ({ id, name, stocks }) => {
  const [watchlistName, setWatchlistName] = useState("");
  useEffect(() => setWatchlistName(name), [name]);
  const handleSave = useCallback(() => {
    return updateWatchlist({ id, name, stocks });
  }, [id, name, stocks]);
  const handleRemoveStock = useCallback(() => {}, []);
  const handleAddStock = useCallback(() => {}, []);
  return (
    <div>
      <h2>{watchlistName}</h2>
      <Button onClick={handleRemoveStock}>Add</Button>
      <Button onClick={handleAddStock}>Add</Button>
      <Button onClick={handleSave}>Save</Button>
    </div>
  );
};

Watchlist.propTypes = {
  id: PropTypes.number,
  name: PropTypes.string,
  stocks: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number.isRequired,
      symbol: PropTypes.string.isRequired,
    })
  ),
};

Watchlist.defaultProps = {
  id: undefined,
  name: "",
  stocks: [],
};

export default Watchlist;

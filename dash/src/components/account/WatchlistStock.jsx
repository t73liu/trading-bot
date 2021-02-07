import React, { useCallback } from "react";
import { ListItem, ListItemIcon, ListItemText } from "@material-ui/core";
import PropTypes from "prop-types";
import { useDispatch, useSelector } from "react-redux";
import { Link } from "react-router-dom";
import { selectStockByID } from "../../state/stocks";
import DeleteButton from "../common/DeleteButton";
import { unwatchStock } from "../../state/watchlists";

const WatchlistStock = ({ watchlistID, stockID }) => {
  const dispatch = useDispatch();
  const stock = useSelector((state) => selectStockByID(state, stockID));
  const handleUnwatchStock = useCallback(() => {
    dispatch(
      unwatchStock({
        watchlistID,
        stockID,
      })
    );
  }, [dispatch, watchlistID, stockID]);
  return (
    <>
      <ListItem>
        <ListItemIcon>
          <DeleteButton onClick={handleUnwatchStock} />
        </ListItemIcon>
        <ListItemText>
          <Link to={`/stocks/${stock?.symbol}`}>
            {stock?.symbol} - {stock?.company}
          </Link>
        </ListItemText>
      </ListItem>
    </>
  );
};

WatchlistStock.propTypes = {
  watchlistID: PropTypes.number.isRequired,
  stockID: PropTypes.number.isRequired,
};

export default WatchlistStock;

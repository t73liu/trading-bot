import React, { useCallback, useState } from "react";
import {
  Accordion,
  AccordionActions,
  AccordionDetails,
  AccordionSummary,
  Button,
  Divider,
  IconButton,
  List,
  ListItemIcon,
  ListItem,
  TextField,
} from "@material-ui/core";
import PropTypes from "prop-types";
import { Add, ExpandMore } from "@material-ui/icons";
import { useDispatch } from "react-redux";
import WatchlistStock from "./WatchlistStock";
import StockLookup from "../common/StockLookup";
import {
  deleteWatchlistThunk,
  updateWatchlistThunk,
  watchStock,
} from "../../state/watchlists";
import { stopEvent } from "../../utils/function";

const Watchlist = ({ id, name: initialName, stockIDs }) => {
  const [name, setName] = useState(initialName);
  const [isAdding, setIsAdding] = useState(false);
  const toggleIsAdding = useCallback(() => {
    setIsAdding((prevState) => !prevState);
  }, [setIsAdding]);
  const dispatch = useDispatch();
  const handleWatchStock = useCallback(
    (e, option) => {
      if (option) {
        dispatch(
          watchStock({
            watchlistID: id,
            stockID: option.id,
          })
        );
      }
    },
    [id, dispatch]
  );
  const handleDelete = useCallback(() => {
    dispatch(deleteWatchlistThunk(id));
  }, [id, dispatch]);
  const handleUpdate = useCallback(() => {
    dispatch(updateWatchlistThunk({ id, name, stockIDs }));
  }, [id, name, stockIDs, dispatch]);
  const handleRename = useCallback(
    (e) => {
      setName(e.target.value);
    },
    [setName]
  );
  return (
    <Accordion key={id} defaultExpanded>
      <AccordionSummary expandIcon={<ExpandMore />}>
        <TextField
          value={name}
          onClick={stopEvent}
          onFocus={stopEvent}
          onChange={handleRename}
        />
      </AccordionSummary>
      <AccordionDetails>
        <List dense>
          {stockIDs?.map((stockID) => (
            <WatchlistStock key={stockID} watchlistID={id} stockID={stockID} />
          ))}
          <ListItem>
            <ListItemIcon>
              <IconButton
                color={isAdding ? "secondary" : "default"}
                onClick={toggleIsAdding}
              >
                <Add />
              </IconButton>
            </ListItemIcon>
            {isAdding && <StockLookup handleStockClick={handleWatchStock} />}
          </ListItem>
        </List>
      </AccordionDetails>
      <Divider />
      <AccordionActions>
        <Button size="small" color="secondary" onClick={handleDelete}>
          Delete
        </Button>
        <Button size="small" color="primary" onClick={handleUpdate}>
          Save
        </Button>
      </AccordionActions>
    </Accordion>
  );
};

Watchlist.propTypes = {
  id: PropTypes.number.isRequired,
  name: PropTypes.string.isRequired,
  stockIDs: PropTypes.arrayOf(PropTypes.number.isRequired).isRequired,
};

export default Watchlist;

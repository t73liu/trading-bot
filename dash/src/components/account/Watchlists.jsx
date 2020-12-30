import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  IconButton,
  List,
  ListItem,
  ListItemSecondaryAction,
  ListItemText,
  Typography,
} from "@material-ui/core";
import { Delete } from "@material-ui/icons";
import {
  fetchWatchlistsThunk,
  selectAllWatchlists,
} from "../../state/watchlists";
import { useTitleContext } from "../../state/title-context";

const Watchlists = () => {
  const { setTitle } = useTitleContext();
  useEffect(() => setTitle("Watchlists"), [setTitle]);
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchWatchlistsThunk());
  }, [dispatch]);
  const watchlists = useSelector(selectAllWatchlists);
  return (
    <div>
      {watchlists?.length > 0 && (
        <Accordion>
          {/* Accordion doesn't accept a Fragment as a child. */}
          {watchlists.map(({ id, name, stockIDs }) => [
            <AccordionSummary key={id}>
              <Typography>{name}</Typography>
            </AccordionSummary>,
            <AccordionDetails key={id}>
              <List>
                {stockIDs?.map((stockID) => (
                  <ListItem key={stockID}>
                    <ListItemText primary={stockID} />
                    <ListItemSecondaryAction>
                      <IconButton>
                        <Delete />
                      </IconButton>
                    </ListItemSecondaryAction>
                  </ListItem>
                ))}
              </List>
            </AccordionDetails>,
          ])}
        </Accordion>
      )}
    </div>
  );
};

export default Watchlists;

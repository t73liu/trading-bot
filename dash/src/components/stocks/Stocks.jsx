import React, { useEffect, useState } from "react";
import { Grid } from "@material-ui/core";
import { useTitleContext } from "../../state/title-context";
import { fetchGapStocks } from "../../data/stocks";
import Snapshots from "./Snapshots";

const Stocks = () => {
  const { setTitle } = useTitleContext();
  const [gapStocks, setGapStocks] = useState({});
  useEffect(() => setTitle("Stocks"), [setTitle]);
  useEffect(() => {
    fetchGapStocks().then((response) => {
      setGapStocks(response);
    });
  }, [setGapStocks]);
  return (
    <div>
      <h2>Stocks</h2>
      <Grid container spacing={1}>
        <Grid item xs>
          <h2>Gap Ups</h2>
          <Snapshots snapshots={gapStocks.gapUp} />
        </Grid>
        <Grid item xs>
          <h2>Gap Downs</h2>
          <Snapshots snapshots={gapStocks.gapDown} />
        </Grid>
      </Grid>
    </div>
  );
};

export default Stocks;

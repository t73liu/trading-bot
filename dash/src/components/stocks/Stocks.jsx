import React, { useEffect, useState } from "react";
import { CircularProgress, Grid } from "@material-ui/core";
import { useTitleContext } from "../../state/title-context";

const Stocks = () => {
  const { setTitle } = useTitleContext();
  const [isLoading] = useState(false);
  useEffect(() => setTitle("Stocks"), [setTitle]);
  if (isLoading) return <CircularProgress />;
  return (
    <div>
      <h2>Stocks</h2>
      <Grid container spacing={1}>
        <Grid item xs>
          <h2>Gap Ups</h2>
        </Grid>
        <Grid item xs>
          <h2>Gap Downs</h2>
        </Grid>
      </Grid>
    </div>
  );
};

export default Stocks;

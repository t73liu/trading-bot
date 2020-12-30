import React, { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import { CircularProgress, Grid } from "@material-ui/core";
import { useTitleContext } from "../../state/title-context";
import EconomicCalendar from "./EconomicCalendar";
import CandlestickChart from "../common/CandlestickChart";
import { fetchStockCharts } from "../../data/stocks";
import { getCandleSize, getShowExtendedHours } from "../../state/settings";

const Market = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [spyCharts, setSPYCharts] = useState({});
  const [qqqCharts, setQQQCharts] = useState({});
  const [vxxCharts, setVXXCharts] = useState({});
  const candleSize = useSelector(getCandleSize);
  const showExtendedHours = useSelector(getShowExtendedHours);
  const { setTitle } = useTitleContext();
  useEffect(() => setTitle("Market"), [setTitle]);
  useEffect(() => {
    setIsLoading(true);
    Promise.all([
      fetchStockCharts("SPY", candleSize, showExtendedHours),
      fetchStockCharts("QQQ", candleSize, showExtendedHours),
      fetchStockCharts("VXX", candleSize, showExtendedHours),
    ]).then(([spyJSON, qqqJSON, vxxJSON]) => {
      if (!(spyJSON instanceof Error)) {
        setSPYCharts(spyJSON);
      }
      if (!(qqqJSON instanceof Error)) {
        setQQQCharts(qqqJSON);
      }
      if (!(vxxJSON instanceof Error)) {
        setVXXCharts(vxxJSON);
      }
      setIsLoading(false);
    });
  }, [candleSize, showExtendedHours]);
  return (
    <div>
      <h2>Market</h2>
      <Grid container spacing={1}>
        <Grid item xs>
          <h2>SPY</h2>
          {isLoading ? (
            <CircularProgress />
          ) : (
            <CandlestickChart data={spyCharts} />
          )}
        </Grid>
        <Grid item xs>
          <h2>QQQ</h2>
          {isLoading ? (
            <CircularProgress />
          ) : (
            <CandlestickChart data={qqqCharts} />
          )}
        </Grid>
        <Grid item xs>
          <h2>VXX</h2>
          {isLoading ? (
            <CircularProgress />
          ) : (
            <CandlestickChart data={vxxCharts} />
          )}
        </Grid>
      </Grid>
      <EconomicCalendar />
    </div>
  );
};

export default Market;

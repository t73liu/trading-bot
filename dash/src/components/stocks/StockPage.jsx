import React, { useCallback, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Alert } from "@material-ui/lab";
import { Button, CircularProgress, Grid } from "@material-ui/core";
import { useSelector } from "react-redux";
import { useTitleContext } from "../../state/title-context";
import CandlestickChart from "../common/CandlestickChart";
import { fetchStockCharts, fetchStockInfo } from "../../data/stocks";
import Articles from "./Articles";
import StockInfo from "./StockInfo";
import { getCandleSize, getShowExtendedHours } from "../../state/account";

const StockPage = () => {
  const candleSize = useSelector(getCandleSize);
  const showExtendedHours = useSelector(getShowExtendedHours);
  const [isLoading, setIsLoading] = useState(false);
  const [info, setInfo] = useState({});
  const [charts, setCharts] = useState({});
  const [errors, setErrors] = useState(undefined);
  const { symbol } = useParams();
  const { setTitle } = useTitleContext();
  useEffect(() => setTitle(`Stock: ${symbol}`), [setTitle, symbol]);
  useEffect(() => {
    if (symbol) {
      setIsLoading(true);
      Promise.all([
        fetchStockInfo(symbol),
        fetchStockCharts(symbol, candleSize, showExtendedHours),
      ]).then(([infoJSON, chartsJSON]) => {
        let apiErrors = [];
        if (infoJSON instanceof Error) {
          apiErrors = [...apiErrors, infoJSON.message];
        } else {
          setInfo(infoJSON);
        }
        if (chartsJSON instanceof Error) {
          apiErrors = [...apiErrors, chartsJSON.message];
        } else {
          setCharts(chartsJSON);
        }
        setErrors(apiErrors);
        setIsLoading(false);
      });
    }
  }, [
    setInfo,
    setCharts,
    setErrors,
    setIsLoading,
    symbol,
    showExtendedHours,
    candleSize,
  ]);
  const handleRefresh = useCallback(() => {
    setIsLoading(true);
    fetchStockCharts(symbol, candleSize, showExtendedHours).then((response) => {
      setIsLoading(false);
      if (response instanceof Error) {
        setErrors([response.message]);
      } else {
        setErrors([]);
        setCharts(response);
      }
    });
  }, [setCharts, symbol, showExtendedHours, candleSize]);
  if (isLoading) return <CircularProgress />;
  return (
    <>
      {errors?.map((msg) => (
        <Alert key={msg} severity="error">
          {msg}
        </Alert>
      ))}
      <h2>
        {info.company || "N/A"} ({symbol})
      </h2>
      <Grid container spacing={1}>
        <Grid item xs>
          <StockInfo
            symbol={symbol}
            info={info}
            lastCandlePrice={
              charts.candles?.[charts.candles?.length - 1]?.close
            }
            currentVolume={charts.currentVolume}
          />
        </Grid>
        <Grid item xs={8}>
          <div>
            <Button
              style={{ float: "right" }}
              variant="contained"
              color="primary"
              onClick={handleRefresh}
            >
              Refresh
            </Button>
          </div>
          <CandlestickChart data={charts} />
        </Grid>
      </Grid>
      <h2>News</h2>
      <Articles articles={info.news} />
    </>
  );
};

export default StockPage;

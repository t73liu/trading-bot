import React, { useCallback, useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Alert } from "@material-ui/lab";
import {
  Button,
  CircularProgress,
  FormControl,
  FormControlLabel,
  Grid,
  FormHelperText,
  MenuItem,
  Select,
  Switch,
} from "@material-ui/core";
import { useTitleContext } from "../../state/title-context";
import CandlestickChart from "../common/CandlestickChart";
import { fetchStockCharts, fetchStockInfo } from "../../data/stocks";
import Articles from "./Articles";
import StockInfo from "./StockInfo";

const StockPage = () => {
  const [candleSize, setCandleSize] = useState("1min");
  const handleCandleSizeChange = useCallback(
    (e) => {
      setCandleSize(e.target.value);
    },
    [setCandleSize]
  );
  const [isLoading, setIsLoading] = useState(false);
  const [showExtendedHours, setShowExtendedHours] = useState(false);
  const toggleShowExtendedHours = useCallback(() => {
    setShowExtendedHours((prevState) => !prevState);
  }, [setShowExtendedHours]);
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
            <FormControlLabel
              control={
                <Switch
                  name="extended"
                  checked={showExtendedHours}
                  onChange={toggleShowExtendedHours}
                />
              }
              label="Show Extended Hours"
            />
            <FormControl>
              <Select value={candleSize} onChange={handleCandleSizeChange}>
                <MenuItem value="1min">1 minute</MenuItem>
                <MenuItem value="3min">3 minute</MenuItem>
                <MenuItem value="5min">5 minute</MenuItem>
                <MenuItem value="10min">10 minute</MenuItem>
                <MenuItem value="30min">30 minute</MenuItem>
                <MenuItem value="1hour">1 hour</MenuItem>
              </Select>
              <FormHelperText>Candle Size</FormHelperText>
            </FormControl>
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

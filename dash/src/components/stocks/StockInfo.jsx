import React from "react";
import { useParams } from "react-router-dom";
import { useTitleContext } from "../../state/title-context";
import CandlestickChart from "../common/CandlestickChart";

const StockInfo = () => {
  const { symbol } = useParams();
  const { setTitle } = useTitleContext();
  setTitle(`Stock: ${symbol}`);
  return (
    <>
      <h2>Stock: {symbol}</h2>
      <CandlestickChart />
    </>
  );
};

export default StockInfo;

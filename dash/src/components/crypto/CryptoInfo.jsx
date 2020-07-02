import React from "react";
import { useParams } from "react-router-dom";
import { useTitleContext } from "../../state/title-context";

const CryptoInfo = () => {
  const { symbol } = useParams();
  const { setTitle } = useTitleContext();
  setTitle(`Crypto: ${symbol}`);
  return (
    <>
      <h2>Crypto: {symbol}</h2>
    </>
  );
};

export default CryptoInfo;

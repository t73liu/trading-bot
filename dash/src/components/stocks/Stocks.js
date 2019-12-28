import React from "react";
import { useTitleContext } from "../../state/title-context";

const Stocks = () => {
  const { setTitle } = useTitleContext();
  setTitle("Stocks");
  return (
    <div>
      <h2>Stocks</h2>
    </div>
  );
};

export default Stocks;

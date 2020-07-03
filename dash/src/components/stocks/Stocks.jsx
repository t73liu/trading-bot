import React, { useEffect } from "react";
import { useTitleContext } from "../../state/title-context";

const Stocks = () => {
  const { setTitle } = useTitleContext();
  useEffect(() => setTitle("Stocks"), [setTitle]);
  return (
    <div>
      <h2>Stocks</h2>
    </div>
  );
};

export default Stocks;

import React from "react";
import { useTitleContext } from "../../state/title-context";

const Crypto = () => {
  const { setTitle } = useTitleContext();
  setTitle("Crypto");
  return (
    <div>
      <h2>Crypto</h2>
    </div>
  );
};

export default Crypto;

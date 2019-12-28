import React from "react";
import { useTitleContext } from "../../state/title-context";

const Overview = () => {
  const { setTitle } = useTitleContext();
  setTitle("Overview");
  return (
    <div>
      <h2>Overview</h2>
    </div>
  );
};

export default Overview;

import React from "react";
import { useTitleContext } from "../state/title-context";

const NotFound = () => {
  const { setTitle } = useTitleContext();
  setTitle("Not Found");
  return (
    <div>
      <h2>Not Found</h2>
    </div>
  );
};

export default NotFound;

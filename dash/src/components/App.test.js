import React from "react";
import { render } from "@testing-library/react";
import App from "./App";

test("renders Overview route", () => {
  const { getByText } = render(<App />);
  const header = getByText(
    (content, element) => content === "Overview" && element.nodeName === "H6"
  );
  expect(header).toBeInTheDocument();
});

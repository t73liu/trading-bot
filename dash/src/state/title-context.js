import { useState } from "react";
import constate from "constate";

function useTitle() {
  const [title, setTitle] = useState("No Title");
  return { title, setTitle };
}

const [TitleProvider, useTitleContext] = constate(useTitle);

export { TitleProvider, useTitleContext };

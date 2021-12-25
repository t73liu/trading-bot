import { createTheme, ThemeProvider } from "@mui/material/styles";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Home from "./Home";
import Layout from "./Layout";
import Login from "./Login";
import NotFound from "./NotFound";
import NavigationScroll from "./common/NavigationScroll";

const theme = createTheme();

const App = (): JSX.Element => (
  <BrowserRouter>
    <NavigationScroll>
      <ThemeProvider theme={theme}>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Layout />}>
            <Route index element={<Home />} />
            <Route path="stocks" element={<Home />} />
            <Route path="watchlists" element={<Home />} />
            <Route path="*" element={<NotFound />} />
          </Route>
        </Routes>
      </ThemeProvider>
    </NavigationScroll>
  </BrowserRouter>
);

export default App;

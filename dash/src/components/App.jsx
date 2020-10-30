import React, { useEffect } from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import { HelmetProvider } from "react-helmet-async";
import { CssBaseline } from "@material-ui/core";
import { useDispatch } from "react-redux";
import Market from "./market/Market";
import Stocks from "./stocks/Stocks";
import StockPage from "./stocks/StockPage";
import Layout from "./Layout";
import NotFound from "./NotFound";
import { TitleProvider } from "../state/title-context";
import Watchlists from "./account/Watchlists";
import { fetchStocksThunk } from "../state/stocks";

const App = () => {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(fetchStocksThunk());
  }, [dispatch]);
  return (
    <HelmetProvider>
      <BrowserRouter>
        <TitleProvider>
          <CssBaseline />
          <Layout>
            <Switch>
              <Route exact path="/stocks/:symbol">
                <StockPage />
              </Route>
              <Route exact path="/stocks">
                <Stocks />
              </Route>
              <Route exact path="/watchlists">
                <Watchlists />
              </Route>
              <Route exact path="/">
                <Market />
              </Route>
              <Route>
                <NotFound />
              </Route>
            </Switch>
          </Layout>
        </TitleProvider>
      </BrowserRouter>
    </HelmetProvider>
  );
};

export default App;

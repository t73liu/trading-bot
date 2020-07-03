import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import { HelmetProvider } from "react-helmet-async";
import { CssBaseline } from "@material-ui/core";
import Overview from "./account/Overview";
import Stocks from "./stocks/Stocks";
import StockInfo from "./stocks/StockInfo";
import Layout from "./Layout";
import NotFound from "./NotFound";
import { TitleProvider } from "../state/title-context";
import Watchlists from "./account/Watchlists";

const App = () => (
  <HelmetProvider>
    <BrowserRouter>
      <TitleProvider>
        <CssBaseline />
        <Layout>
          <Switch>
            <Route exact path="/stocks/:symbol">
              <StockInfo />
            </Route>
            <Route exact path="/stocks">
              <Stocks />
            </Route>
            <Route exact path="/watchlists">
              <Watchlists />
            </Route>
            <Route exact path="/">
              <Overview />
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

export default App;

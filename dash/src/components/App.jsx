import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import { HelmetProvider } from "react-helmet-async";
import { CssBaseline } from "@material-ui/core";
import Overview from "./overview/Overview";
import Stocks from "./stocks/Stocks";
import StockInfo from "./stocks/StockInfo";
import Crypto from "./crypto/Crypto";
import CryptoInfo from "./crypto/CryptoInfo";
import Layout from "./Layout";
import NotFound from "./NotFound";
import { TitleProvider } from "../state/title-context";

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
            <Route exact path="/crypto/:symbol">
              <CryptoInfo />
            </Route>
            <Route exact path="/crypto">
              <Crypto />
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

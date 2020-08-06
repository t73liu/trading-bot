import React, { useCallback, useEffect, useState } from "react";
import { Link, useParams } from "react-router-dom";
import { Alert } from "@material-ui/lab";
import {
  Button,
  Chip,
  CircularProgress,
  FormControlLabel,
  Grid,
  Switch,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@material-ui/core";
import PropTypes from "prop-types";
import dayjs from "dayjs";
import { useTitleContext } from "../../state/title-context";
import CandlestickChart from "../common/CandlestickChart";
import {
  fetchStockCandles,
  fetchStockInfo,
  fetchStockNews,
} from "../../data/stocks";
import ExternalLink from "../common/ExternalLink";
import { formatAsCurrency } from "../../utils/number";

const OverflowDiv = ({ children }) => (
  <div style={{ maxHeight: "100px", overflow: "auto" }}>{children}</div>
);

OverflowDiv.propTypes = {
  children: PropTypes.node.isRequired,
};

const formatCandle = ({ openedAt, ...candle }) => ({
  x: openedAt,
  ...candle,
});

const StockInfo = () => {
  const [isLoading, setIsLoading] = useState(false);
  const [showExtendedHours, setShowExtendedHours] = useState(false);
  const toggleShowExtendedHours = useCallback(() => {
    setShowExtendedHours((prevState) => !prevState);
  }, [setShowExtendedHours]);
  const [info, setInfo] = useState({});
  const [, setNews] = useState([]);
  const [candles, setCandles] = useState([]);
  const [errors, setErrors] = useState(undefined);
  const { symbol } = useParams();
  const { setTitle } = useTitleContext();
  useEffect(() => setTitle(`Stock: ${symbol}`), [setTitle, symbol]);
  useEffect(() => {
    if (symbol) {
      setIsLoading(true);
      Promise.all([
        fetchStockInfo(symbol),
        fetchStockNews(symbol),
        fetchStockCandles(symbol, showExtendedHours),
      ]).then(([infoJSON, newsJSON, candlesJSON]) => {
        let apiErrors = [];
        if (infoJSON instanceof Error) {
          apiErrors = [...apiErrors, infoJSON.message];
        } else {
          setInfo(infoJSON);
        }
        if (newsJSON instanceof Error) {
          apiErrors = [...apiErrors, newsJSON.message];
        } else {
          setNews(newsJSON);
        }
        if (candlesJSON instanceof Error) {
          apiErrors = [...apiErrors, candlesJSON.message];
        } else {
          setCandles(candlesJSON?.map(formatCandle));
        }
        setErrors(apiErrors);
        setIsLoading(false);
      });
    }
  }, [
    setInfo,
    setNews,
    setErrors,
    setCandles,
    setIsLoading,
    symbol,
    showExtendedHours,
  ]);
  const handleRefresh = useCallback(() => {
    setIsLoading(true);
    fetchStockCandles(symbol, showExtendedHours).then((response) => {
      setIsLoading(false);
      if (response instanceof Error) {
        setErrors([response.message]);
      } else {
        setCandles(response?.map(formatCandle));
      }
    });
  }, [symbol, showExtendedHours, setCandles]);
  if (isLoading) return <CircularProgress />;
  return (
    <>
      {errors?.map((msg) => (
        <Alert key={msg} severity="error">
          {msg}
        </Alert>
      ))}
      <h2>
        {info.company || "N/A"} ({symbol})
      </h2>
      <Grid container spacing={1}>
        <Grid item xs>
          <TableContainer>
            <Table>
              <TableBody>
                <TableRow>
                  <TableCell>Market Cap</TableCell>
                  <TableCell>{formatAsCurrency(info.marketCap)}</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Country</TableCell>
                  <TableCell>{info.country}</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Sector</TableCell>
                  <TableCell>{info.sector}</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Industry</TableCell>
                  <TableCell>{info.industry}</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Description</TableCell>
                  <TableCell>
                    <OverflowDiv>{info.description}</OverflowDiv>
                  </TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Website</TableCell>
                  <TableCell>
                    <ExternalLink url={info.website}>
                      {info.website}
                    </ExternalLink>
                  </TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Similar</TableCell>
                  <TableCell>
                    {info.similarStocks?.sort().map((s) => (
                      <Link to={`/stocks/${s}`}>
                        <Chip label={s} clickable color="primary" />
                      </Link>
                    ))}
                  </TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Useful Links</TableCell>
                  <TableCell>
                    <ExternalLink
                      url={`https://finance.yahoo.com/quote/${symbol}/`}
                    >
                      <Chip label="Yahoo" clickable />
                    </ExternalLink>
                    <ExternalLink
                      url={`https://seekingalpha.com/symbol/${symbol}/`}
                    >
                      <Chip label="SeekingAlpha" clickable />
                    </ExternalLink>
                    <ExternalLink
                      url={`https://stocktwits.com/symbol/${symbol}/`}
                    >
                      <Chip label="Stocktwits" clickable />
                    </ExternalLink>
                  </TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Marginable</TableCell>
                  <TableCell>{info.marginable ? "YES" : "NO"}</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell>Shortable</TableCell>
                  <TableCell>{info.shortable ? "YES" : "NO"}</TableCell>
                </TableRow>
              </TableBody>
            </Table>
          </TableContainer>
        </Grid>
        <Grid item xs={8}>
          <div>
            <FormControlLabel
              control={
                <Switch
                  name="extended"
                  checked={showExtendedHours}
                  onChange={toggleShowExtendedHours}
                />
              }
              label="Show Extended Hours"
            />
            <Button variant="contained" color="primary" onClick={handleRefresh}>
              Refresh
            </Button>
          </div>
          <CandlestickChart data={candles} />
        </Grid>
      </Grid>
      <h2>News</h2>
      <TableContainer>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Title</TableCell>
              <TableCell>Description</TableCell>
              <TableCell>Published At</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {info.news?.map((article) => (
              <TableRow key={article.url}>
                <TableCell>
                  <ExternalLink url={article.url}>{article.title}</ExternalLink>
                </TableCell>
                <TableCell>
                  <OverflowDiv>{article.summary}...</OverflowDiv>
                </TableCell>
                <TableCell>
                  {dayjs(article.publishedAt).format("YYYY-MM-DD HH:mm:ss")}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
};

export default StockInfo;

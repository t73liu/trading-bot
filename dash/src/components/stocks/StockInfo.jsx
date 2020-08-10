import React from "react";
import {
  Chip,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableRow,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import PropTypes from "prop-types";
import { formatAsCurrency, formatWithCommas } from "../../utils/number";
import ExternalLink from "../common/ExternalLink";
import OverflowDiv from "../common/OverflowDiv";

const StockInfo = ({ info, symbol, currentVolume }) => (
  <TableContainer>
    <Table>
      <TableBody>
        <TableRow>
          <TableCell>Price</TableCell>
          <TableCell>{formatAsCurrency(info.price)}</TableCell>
        </TableRow>
        <TableRow>
          <TableCell>Market Cap</TableCell>
          <TableCell>{formatAsCurrency(info.marketCap)}</TableCell>
        </TableRow>
        <TableRow>
          <TableCell>Current Volume</TableCell>
          <TableCell>{formatWithCommas(currentVolume)}</TableCell>
        </TableRow>
        <TableRow>
          <TableCell>Average Volume</TableCell>
          <TableCell>{formatWithCommas(info.averageVolume)}</TableCell>
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
            {info.website && (
              <ExternalLink url={info.website}>{info.website}</ExternalLink>
            )}
          </TableCell>
        </TableRow>
        <TableRow>
          <TableCell>Similar</TableCell>
          <TableCell>
            {info.similarStocks?.sort().map((s) => (
              <Link key={s} to={`/stocks/${s}`}>
                <Chip label={s} clickable color="primary" />
              </Link>
            ))}
          </TableCell>
        </TableRow>
        <TableRow>
          <TableCell>Useful Links</TableCell>
          <TableCell>
            <ExternalLink url={`https://finance.yahoo.com/quote/${symbol}/`}>
              <Chip label="Yahoo" clickable />
            </ExternalLink>
            <ExternalLink url={`https://seekingalpha.com/symbol/${symbol}/`}>
              <Chip label="SeekingAlpha" clickable />
            </ExternalLink>
            <ExternalLink url={`https://stocktwits.com/symbol/${symbol}/`}>
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
);

StockInfo.propTypes = {
  symbol: PropTypes.string.isRequired,
  currentVolume: PropTypes.number,
  info: PropTypes.shape({
    price: PropTypes.number,
    industry: PropTypes.string,
    sector: PropTypes.string,
    website: PropTypes.string,
    country: PropTypes.string,
    description: PropTypes.string,
    averageVolume: PropTypes.number,
    marketCap: PropTypes.number,
    similarStocks: PropTypes.arrayOf(PropTypes.string.isRequired),
    shortable: PropTypes.bool.isRequired,
    marginable: PropTypes.bool.isRequired,
  }).isRequired,
};

StockInfo.defaultProps = {
  currentVolume: 0,
};

export default StockInfo;

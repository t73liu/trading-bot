import React from "react";
import PropTypes from "prop-types";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@material-ui/core";
import { Link } from "react-router-dom";
import {
  formatAsCurrency,
  formatAsPercent,
  formatWithCommas,
} from "../../utils/number";
import { formatDateTime } from "../../utils/time";

const Snapshots = ({ snapshots }) => (
  <TableContainer>
    <Table>
      <TableHead>
        <TableRow>
          <TableCell>Stock</TableCell>
          <TableCell>Change</TableCell>
          <TableCell>Previous Close</TableCell>
          <TableCell>Previous Volume</TableCell>
          <TableCell>Last Updated</TableCell>
        </TableRow>
      </TableHead>
      <TableBody>
        {snapshots?.map((snapshot) => (
          <TableRow key={snapshot.symbol}>
            <TableCell>
              <Link to={`/stocks/${snapshot.symbol}`}>
                {snapshot.company} ({snapshot.symbol})
              </Link>
            </TableCell>
            <TableCell style={{ color: snapshot.change > 0 ? "green" : "red" }}>
              {snapshot.change > 0 && "+"}
              {formatAsCurrency(snapshot.change)} (
              {formatAsPercent(snapshot.changePercent / 100)})
            </TableCell>
            <TableCell>{formatAsCurrency(snapshot.previousClose)}</TableCell>
            <TableCell>{formatWithCommas(snapshot.previousVolume)}</TableCell>
            <TableCell>{formatDateTime(snapshot.updatedAt)}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  </TableContainer>
);

Snapshots.propTypes = {
  snapshots: PropTypes.arrayOf(
    PropTypes.shape({
      symbol: PropTypes.string.isRequired,
      company: PropTypes.string.isRequired,
      change: PropTypes.number.isRequired,
      changePercent: PropTypes.number.isRequired,
      previousClose: PropTypes.number.isRequired,
      previousVolume: PropTypes.number.isRequired,
      updatedAt: PropTypes.string.isRequired,
    })
  ),
};

Snapshots.defaultProps = {
  snapshots: [],
};

export default Snapshots;

import React from "react";
import { Autocomplete, createFilterOptions } from "@material-ui/lab";
import { TextField } from "@material-ui/core";
import { useSelector } from "react-redux";
import PropTypes from "prop-types";
import { selectAllStocks } from "../../state/stocks";

const filterOptions = createFilterOptions({
  limit: 50,
});

const getStockLabel = (stock) => `${stock.symbol} - ${stock.company}`;

const StockLookup = ({ handleStockClick }) => {
  const stocks = useSelector(selectAllStocks);
  return (
    <Autocomplete
      onChange={handleStockClick}
      filterOptions={filterOptions}
      options={stocks}
      getOptionLabel={getStockLabel}
      renderInput={(params) => (
        <TextField
          {...params}
          margin="normal"
          variant="outlined"
          placeholder="Symbol"
          style={{
            backgroundColor: "white",
            width: 300,
          }}
        />
      )}
    />
  );
};

StockLookup.propTypes = {
  handleStockClick: PropTypes.func.isRequired,
};

export default StockLookup;

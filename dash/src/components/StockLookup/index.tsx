import {
  createFilterOptions,
  Autocomplete,
  TextField,
  Skeleton,
} from "@mui/material";
import { useMemo } from "react";

import { useGetStocksQuery } from "../../state/stocks";
import { Stock } from "../../types/stocks";
import { noop } from "../../utils/function";

const filterOptions = createFilterOptions({
  limit: 50,
});

const StockLookup = (): JSX.Element => {
  const { data: stocks, isLoading } = useGetStocksQuery();
  const options = useMemo(() => {
    if (!stocks) return [];
    return stocks.map((stock: Stock) => ({
      id: stock.symbol,
      label: `${stock.symbol} - ${stock.company}`,
    }));
  }, [stocks]);
  if (isLoading) {
    return <Skeleton />;
  }
  return (
    <Autocomplete
      onChange={noop}
      filterOptions={filterOptions}
      options={options}
      renderInput={(params): JSX.Element => (
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

export default StockLookup;

import React from "react";
import { VictoryAxis, VictoryCandlestick, VictoryChart } from "victory";
import PropTypes from "prop-types";
import dayjs from "dayjs";

const formatDateStr = (dateStr) => {
  if (typeof dateStr === "string") {
    return dayjs(dateStr).format("HH:mm");
  }
  return "N/A";
};

const CandlestickChart = ({ data, className }) => {
  return (
    <div className={className}>
      <VictoryChart scale={{ x: "time" }}>
        <VictoryAxis
          label="Time (min)"
          tickCount={5}
          tickFormat={formatDateStr}
        />
        <VictoryAxis dependentAxis />
        <VictoryCandlestick
          candleColors={{ positive: "#90EE90", negative: "#8B0000" }}
          candleRatio={0.5}
          data={data}
        />
      </VictoryChart>
    </div>
  );
};

CandlestickChart.propTypes = {
  className: PropTypes.string,
  data: PropTypes.arrayOf(
    PropTypes.shape({
      x: PropTypes.string.isRequired,
      open: PropTypes.number.isRequired,
      close: PropTypes.number.isRequired,
      high: PropTypes.number.isRequired,
      low: PropTypes.number.isRequired,
      volume: PropTypes.number.isRequired,
    })
  ),
};

CandlestickChart.defaultProps = {
  className: undefined,
  data: [],
};

export default CandlestickChart;

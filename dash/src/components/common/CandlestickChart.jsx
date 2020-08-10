import React from "react";
import {
  VictoryAxis,
  VictoryCandlestick,
  VictoryChart,
  VictoryLine,
  VictoryBar,
} from "victory";
import PropTypes from "prop-types";
import dayjs from "dayjs";

const formatDateStr = (dateStr) => {
  if (typeof dateStr === "string") {
    return dayjs(dateStr).format("HH:mm");
  }
  return "N/A";
};

const CandlestickChart = ({ data, className }) => {
  const { candles, vwap, volume } = data;
  let min = 0;
  let max = 0;
  // eslint-disable-next-line no-unused-expressions
  candles?.forEach((c) => {
    if (!max || c.high > max) {
      max = c.high;
    }
    if (!min || c.low < min) {
      min = c.low;
    }
  });
  return (
    <div className={className}>
      <VictoryChart scale={{ x: "time" }}>
        <VictoryAxis
          label="Time (min)"
          tickCount={5}
          tickFormat={formatDateStr}
        />
        <VictoryAxis dependentAxis />
        {candles?.length > 0 && (
          <VictoryCandlestick
            candleColors={{ positive: "#90EE90", negative: "#8B0000" }}
            candleRatio={0.5}
            data={candles}
          />
        )}
        {vwap?.length > 0 && <VictoryLine data={vwap} />}
        {/* TODO figure out multi-line chart */}
        {/* {ema?.length > 0 && <VictoryLine data={ema} />} */}
      </VictoryChart>
      <VictoryChart height={160}>
        {volume?.length > 0 && <VictoryBar data={volume} />}
      </VictoryChart>
    </div>
  );
};

CandlestickChart.propTypes = {
  className: PropTypes.string,
  data: PropTypes.shape({
    candles: PropTypes.arrayOf(
      PropTypes.shape({
        x: PropTypes.string.isRequired,
        open: PropTypes.number.isRequired,
        close: PropTypes.number.isRequired,
        high: PropTypes.number.isRequired,
        low: PropTypes.number.isRequired,
        volume: PropTypes.number.isRequired,
      })
    ),
    vwap: PropTypes.arrayOf(
      PropTypes.shape({
        x: PropTypes.string.isRequired,
        value: PropTypes.number.isRequired,
      })
    ),
    volume: PropTypes.arrayOf(
      PropTypes.shape({
        x: PropTypes.string.isRequired,
        value: PropTypes.number.isRequired,
      })
    ),
    ema: PropTypes.arrayOf(
      PropTypes.shape({
        x: PropTypes.string.isRequired,
        value: PropTypes.number,
      })
    ),
  }),
};

CandlestickChart.defaultProps = {
  className: undefined,
  data: {},
};

export default CandlestickChart;

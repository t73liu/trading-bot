import React from "react";
import {
  VictoryAxis,
  VictoryCandlestick,
  VictoryChart,
  VictoryLine,
  VictoryBar,
  VictoryTooltip,
} from "victory";
import PropTypes from "prop-types";
import dayjs from "dayjs";
import { formatWithCommas, humanizeNumber } from "../../utils/number";

const formatDateStr = (dateStr) => {
  if (typeof dateStr === "string") {
    return dayjs(dateStr).format("HH:mm");
  }
  return "N/A";
};

const TICK_LABEL_SIZE = { tickLabels: { fontSize: 8 } };

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
          style={TICK_LABEL_SIZE}
          tickCount={8}
          tickFormat={formatDateStr}
        />
        <VictoryAxis dependentAxis style={TICK_LABEL_SIZE} />
        {candles?.length > 0 && (
          <VictoryCandlestick
            labelComponent={<VictoryTooltip />}
            candleColors={{ positive: "#90EE90", negative: "#8B0000" }}
            candleRatio={0.5}
            data={candles?.map((c) => ({ ...c, label: JSON.stringify(c) }))}
            x="openedAt"
          />
        )}
        {vwap?.length > 0 && (
          <VictoryLine
            data={vwap}
            style={{ data: { stroke: "gray", strokeWidth: 0.5 } }}
          />
        )}
        {/* TODO figure out multi-line chart */}
        {/* {ema?.length > 0 && <VictoryLine data={ema} />} */}
      </VictoryChart>
      <VictoryChart height={160} scale={{ x: "time" }}>
        <VictoryAxis
          style={TICK_LABEL_SIZE}
          tickCount={8}
          tickFormat={formatDateStr}
        />
        <VictoryAxis
          dependentAxis
          tickFormat={humanizeNumber}
          style={TICK_LABEL_SIZE}
        />
        {volume?.length > 0 && (
          <VictoryBar
            style={{
              data: {
                fill: ({ datum }) => datum.fill,
              },
            }}
            data={volume?.map((y, i) => {
              const candle = candles[i];
              return {
                x: candle.openedAt,
                y,
                label: formatWithCommas(y),
                fill: candles[i].close > candles[i].open ? "green" : "red",
              };
            })}
            labelComponent={<VictoryTooltip />}
          />
        )}
      </VictoryChart>
    </div>
  );
};

CandlestickChart.propTypes = {
  className: PropTypes.string,
  data: PropTypes.shape({
    candles: PropTypes.arrayOf(
      PropTypes.shape({
        openedAt: PropTypes.string.isRequired,
        open: PropTypes.number.isRequired,
        close: PropTypes.number.isRequired,
        high: PropTypes.number.isRequired,
        low: PropTypes.number.isRequired,
        volume: PropTypes.number.isRequired,
      })
    ),
    vwap: PropTypes.arrayOf(PropTypes.number.isRequired),
    volume: PropTypes.arrayOf(PropTypes.number.isRequired),
    ema: PropTypes.arrayOf(PropTypes.number),
  }),
};

CandlestickChart.defaultProps = {
  className: undefined,
  data: {},
};

export default CandlestickChart;

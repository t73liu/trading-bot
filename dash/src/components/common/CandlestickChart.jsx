import React from "react";
import {
  VictoryAxis,
  VictoryBar,
  VictoryCandlestick,
  VictoryChart,
  VictoryLabel,
  VictoryLegend,
  VictoryLine,
  VictoryScatter,
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
  const { candles, ema, macd, rsi, ttmSqueeze, vwap, volume } = data;
  let min = 0;
  let max = 0;
  candles?.forEach((c) => {
    if (!max || c.high > max) {
      max = c.high;
    }
    if (!min || c.low < min) {
      min = c.low;
    }
  });
  const formatMACD = (val, i) => {
    let fill;
    if (val < 0) {
      fill = "red";
      if (i > 0 && val > macd?.[i - 1]) {
        fill = "#FFCCCB";
      }
    } else {
      fill = "green";
      if (i > 0 && val < macd?.[i - 1]) {
        fill = "#90EE90";
      }
    }
    return {
      openedAt: candles?.[i].openedAt,
      y: val,
      fill,
    };
  };
  return (
    <div className={className}>
      <VictoryChart scale={{ x: "time" }} domain={{ y: [min, max] }}>
        <VictoryAxis
          style={TICK_LABEL_SIZE}
          tickCount={8}
          tickFormat={formatDateStr}
        />
        <VictoryAxis dependentAxis style={TICK_LABEL_SIZE} />
        <VictoryLegend
          x={150}
          y={50}
          orientation="horizontal"
          gutter={20}
          style={{ border: { stroke: "black" }, labels: { fontSize: 8 } }}
          data={[
            { name: "VWAP", symbol: { fill: "#FF7700" } },
            { name: "EMA", symbol: { fill: "#008080" } },
            { name: "Squeeze", symbol: { fill: "purple" } },
          ]}
        />
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
            x="openedAt"
            data={vwap?.map((y, i) => ({ openedAt: candles?.[i].openedAt, y }))}
            style={{ data: { stroke: "#FF7700", strokeWidth: 0.8 } }}
          />
        )}
        {ema?.length > 0 && (
          <VictoryLine
            x="openedAt"
            data={ema?.map((y, i) => ({ openedAt: candles?.[i].openedAt, y }))}
            style={{ data: { stroke: "#003A80", strokeWidth: 0.8 } }}
          />
        )}
        {ttmSqueeze?.length > 0 && (
          <VictoryScatter
            style={{
              data: { fill: "purple" },
            }}
            size={2}
            data={ttmSqueeze?.map((y, i) => {
              const candle = candles[i];
              return {
                x: candle.openedAt,
                y: y ? min : null,
              };
            })}
          />
        )}
      </VictoryChart>
      {macd?.length > 0 && (
        <VictoryChart height={150} scale={{ x: "time" }}>
          <VictoryLabel text="MACD" x={225} y={30} textAnchor="middle" />
          <VictoryAxis
            style={TICK_LABEL_SIZE}
            tickCount={8}
            tickFormat={formatDateStr}
          />
          <VictoryAxis
            dependentAxis
            tickCount={3}
            tickFormat={humanizeNumber}
            style={TICK_LABEL_SIZE}
          />
          <VictoryBar
            x="openedAt"
            style={{
              data: {
                fill: ({ datum }) => datum.fill,
              },
            }}
            data={macd?.map(formatMACD)}
          />
        </VictoryChart>
      )}
      {rsi?.length > 0 && (
        <VictoryChart
          height={150}
          scale={{ x: "time" }}
          domain={{ y: [0, 100] }}
        >
          <VictoryLabel text="RSI" x={225} y={30} textAnchor="middle" />
          <VictoryAxis
            style={TICK_LABEL_SIZE}
            tickCount={8}
            tickFormat={formatDateStr}
          />
          <VictoryAxis
            dependentAxis
            style={TICK_LABEL_SIZE}
            tickFormat={humanizeNumber}
          />
          <VictoryLine
            x="openedAt"
            data={rsi?.map((y, i) => ({ openedAt: candles?.[i].openedAt, y }))}
            style={{ data: { stroke: "blue", strokeWidth: 0.8 } }}
          />
          <VictoryLine
            style={{
              data: { stroke: "red", strokeWidth: 0.3 },
            }}
            y={() => 70}
          />
          <VictoryLine
            style={{
              data: { stroke: "red", strokeWidth: 0.3 },
            }}
            y={() => 30}
          />
        </VictoryChart>
      )}
      {volume?.length > 0 && (
        <VictoryChart height={150} scale={{ x: "time" }}>
          <VictoryLabel text="Volume" x={225} y={30} textAnchor="middle" />
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
          <VictoryBar
            x="openedAt"
            style={{
              data: {
                fill: ({ datum }) => datum.fill,
              },
            }}
            data={volume?.map((y, i) => {
              const candle = candles[i];
              return {
                openedAt: candle.openedAt,
                y,
                label: formatWithCommas(y),
                fill: candles[i].close > candles[i].open ? "green" : "red",
              };
            })}
            labelComponent={<VictoryTooltip />}
          />
        </VictoryChart>
      )}
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
    ema: PropTypes.arrayOf(PropTypes.number),
    macd: PropTypes.arrayOf(PropTypes.number),
    rsi: PropTypes.arrayOf(PropTypes.number),
    ttmSqueeze: PropTypes.arrayOf(PropTypes.bool),
    volume: PropTypes.arrayOf(PropTypes.number.isRequired),
    vwap: PropTypes.arrayOf(PropTypes.number.isRequired),
  }),
};

CandlestickChart.defaultProps = {
  className: undefined,
  data: {},
};

export default CandlestickChart;

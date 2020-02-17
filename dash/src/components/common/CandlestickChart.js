import React, { useState } from "react";
import {
  VictoryChart,
  VictoryCandlestick,
  VictoryAxis,
  VictoryLine,
  VictoryZoomContainer,
} from "victory";

const fakeData = [
  { x: new Date(2016, 6, 1), open: 9, close: 30, high: 56, low: 7 },
  { x: new Date(2016, 6, 2), open: 80, close: 40, high: 120, low: 10 },
  { x: new Date(2016, 6, 3), open: 50, close: 80, high: 90, low: 20 },
  { x: new Date(2016, 6, 4), open: 70, close: 22, high: 70, low: 5 },
];

const fakeDomain = {
  x: [new Date(2016, 6, 1), new Date(2016, 6, 2)],
};

function getCandleVolume(candle) {
  return candle.volume;
}

const CandlestickChart = ({ data = fakeData }) => {
  const [domain, setDomain] = useState(fakeDomain);
  return (
    <div>
      <VictoryChart
        width={600}
        height={500}
        domainPadding={{ x: 25 }}
        scale={{ x: "time" }}
        containerComponent={
          <VictoryZoomContainer
            zoomDimension="x"
            zoomDomain={domain}
            onZoomDomainChange={setDomain}
          />
        }
      >
        <VictoryAxis tickFormat={(t) => `${t.getDate()}/${t.getMonth()}`} />
        <VictoryAxis dependentAxis />
        <VictoryCandlestick data={data} />
      </VictoryChart>
      <VictoryChart
        width={600}
        height={500}
        domainPadding={{ x: 25 }}
        scale={{ x: "time" }}
        containerComponent={
          <VictoryZoomContainer
            zoomDimension="x"
            zoomDomain={domain}
            onZoomDomainChange={setDomain}
          />
        }
      >
        <VictoryLine data={fakeData} y={getCandleVolume} />
      </VictoryChart>
    </div>
  );
};

export default CandlestickChart;

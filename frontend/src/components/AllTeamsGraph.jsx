import ReactECharts from "echarts-for-react";
import { Box } from "@mui/material";
import { useMemo } from "react";

// const SLICE=400; // Limit the number of entries to max num of teams

export default function AllTeamsGraph({ data }) {
  const options = useMemo(() => {
    if (!data) return null;
    
    const teamIds = Object.keys(data);
    const SLICE = teamIds.length;

    const labels = Array(SLICE);
    const pending = Array(SLICE);
    const fail = Array(SLICE);
    const success = Array(SLICE);
    const expired = Array(SLICE);
  
    for (let i = 0; i < SLICE; i++) {
      const teamId = i.toString();
      const teamData = data[teamId];
      const [p = 0, s = 0, f = 0, e = 0] = teamData?.status || [0, 0, 0, 0];
      
      labels[i] = teamId;
      pending[i] = p;
      fail[i] = f;
      success[i] = s;
      expired[i] = e;
    }
  
    return {
      tooltip: {
        trigger: "axis",
        axisPointer: { type: "shadow" },
      },
      grid: {
        top: 10,
        left: "3%",
        right: "4%",
        bottom: "10%",
        containLabel: true,
      },
      xAxis: {
        type: "category",
        data: labels,
        axisLabel: {
          show: labels.length <= 100,
          fontSize: 10,
        },
      },
      yAxis: {
        type: "value",
        splitLine: {
          show: false,
        },
      },
      series: [
        {
          name: "Success",
          type: "bar",
          stack: "total",
          data: success,
          itemStyle: { color: "#00ff84" },
        },
        {
          name: "Failure",
          type: "bar",
          stack: "total",
          data: fail,
          itemStyle: { color: "#ff5040" },
        },
        {
          name: "Expired",
          type: "bar",
          stack: "total",
          data: expired,
          itemStyle: { color: "#ff9900" },
        },
        {
          name: "Pending",
          type: "bar",
          stack: "total",
          data: pending,
          itemStyle: { color: "#00c3ff" },
        },
      ],
    };
  }, [data]);

  if (!options) return null;

  return (
    <Box flexGrow={1}>
      <ReactECharts option={options} style={{ height: 300, width: "100%" }} />
    </Box>
  );
}

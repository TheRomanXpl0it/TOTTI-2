import { PieChart } from "@mui/x-charts/PieChart";
import { Box } from "@mui/material";

export default function TotalPieGraph({ data }) {
  const pieData = [
    { label: "Pending", value: 0, color: "#00c3ff" },
    { label: "Success", value: 0, color: "#00ff84" },
    { label: "Error", value: 0, color: "#ff5040" },
    { label: "Expired", value: 0, color: "#ff9900" },
  ];

  if (data && typeof data === "object") {
    Object.values(data).forEach((entry) => {
      const [pending, success, error, expired] = entry.status || [];
      pieData[0].value += pending || 0;
      pieData[1].value += success || 0;
      pieData[2].value += error || 0;
      pieData[3].value += expired || 0;
    });
  }

  return (
    <Box>
      <PieChart
        series={[
          {
            data: pieData,
            innerRadius: 30,
            outerRadius: 80,
            paddingAngle: 2,
            cornerRadius: 5,
          },
        ]}
        height={300}
      />
    </Box>
  );
}

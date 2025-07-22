import { Box, FormControl, InputLabel, MenuItem, Select } from "@mui/material";
import ManualSubmission from "./ManualSubmission";
import CustomTitle from "./CustomTitle";

function Navbar({
  exploitData,
  selectedExploit,
  setSelectedExploit,
  selectedTime,
  setSelectedTime,
}) {

  const exploitList = ["All", ...Object.keys(exploitData)];

  const timeLists = [
    { label: "Since ever", value: "Since ever" },
    { label: "Last tick", value: 1 },
    { label: "Last 2 ticks", value: 2 },
    { label: "Last 5 ticks", value: 5 },
    { label: "Last 10 ticks", value: 10 },
    { label: "Last 20 ticks", value: 20 },
    { label: "Last 40 ticks", value: 40 },
  ];

  return (
    <Box
      display="flex"
      flexDirection="row"
      alignItems={"start"}
      justifyContent={"space-between"}
      gap={3}
      pb={5}
    >
      <Box display="flex" flexDirection="row" alignItems={"center"} gap={2}>
        <FormControl sx={{ minWidth: 200 }}>
          <InputLabel id="exploit-select-label">Exploit</InputLabel>
          <Select
            labelId="exploit-select-label"
            value={selectedExploit}
            size={"small"}
            label="Exploit"
            onChange={(e) => setSelectedExploit(e.target.value)}
          >
            {exploitList.map((item) => (
              <MenuItem key={item} value={item}>
                {item}
              </MenuItem>
            ))}
          </Select>
        </FormControl>

        <FormControl sx={{ minWidth: 200 }}>
          <InputLabel id="time-select-label">Time</InputLabel>
          <Select
            labelId="time-select-label"
            value={selectedTime}
            size="small"
            label="Time"
            onChange={(e) => setSelectedTime(e.target.value)}
          >
            {timeLists.map(({ label, value }) => (
              <MenuItem key={value} value={value}>
                {label}
              </MenuItem>
            ))}
          </Select>
        </FormControl>
      </Box>
      <Box>
        <CustomTitle />
      </Box>
      <ManualSubmission flexGrow={1} />
    </Box>
  );
}

export default Navbar;

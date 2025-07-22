import "./App.css";
import { Box, Tabs, Tab } from "@mui/material";
import { useEffect, useState } from "react";
import ExploitGraph from "./components/ExploitGraph";
import AllTeamsGraph from "./components/AllTeamsGraph";
import GroupsIcon from "@mui/icons-material/Groups";
import TimelineGraph from "./components/TimelineGraph";
import TotalPieGraph from "./components/TotalPieGraph";
import AccessTimeIcon from "@mui/icons-material/AccessTime";
import Navbar from "./components/Navbar";

function TabPanel({ children, value, index }) {
  return (
    <Box hidden={value !== index} p={2}>
      {value === index && children}
    </Box>
  );
}

function App() {
  const cookies = Object.fromEntries(
    document.cookie.split("; ").map((c) => c.split("="))
  );

  const [exploitData, setExploitData] = useState([]);
  const [teamsData, setTeamsData] = useState([]);
  const [timelineData, setTimelineData] = useState([]);
  const [selectedExploit, setSelectedExploit] = useState(
    cookies.exploit || "All"
  );
  const [selectedTime, setSelectedTime] = useState(
    cookies.round || "Since ever"
  );
  const [tabIndex, setTabIndex] = useState(0);

  useEffect(() => {
    if (selectedExploit === "All") {
      document.cookie =
        "exploit=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    } else {
      document.cookie = `exploit=${selectedExploit}; path=/;`;
    }

    if (selectedTime === "Since ever") {
      document.cookie =
        "round=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    } else {
      document.cookie = `round=${selectedTime}; path=/;`;
    }

    const fetchData = () => {
      fetch("/api/exploit")
        .then((res) => res.json())
        .then((data) => {
          setExploitData(data);
          console.log("Exploit Data:", data);
        })
        .catch((err) => console.error("Error fetching /exploit:", err));

      fetch("/api/timeline")
        .then((res) => res.json())
        .then((data) => {
          setTimelineData(data);
        })
        .catch((err) => console.error("Error fetching /timeline:", err));

      fetch("/api/teams")
        .then((res) => res.json())
        .then((data) => {
          setTeamsData(data);
        })
        .catch((err) => console.error("Error fetching /teams:", err));
    };

    fetchData();
    const interval = setInterval(fetchData, 5000);

    // Cleanup on unmount
    return () => clearInterval(interval);
  }, [selectedExploit, selectedTime]);

  const handleTabChange = (_, newIndex) => {
    setTabIndex(newIndex);
  };

  return (
    <Box>
      <Navbar
        exploitData={exploitData}
        selectedExploit={selectedExploit}
        setSelectedExploit={setSelectedExploit}
        selectedTime={selectedTime}
        setSelectedTime={setSelectedTime}
      />
      <Box display={"flex"} justifyContent="center" alignItems={"center"}>
        <ExploitGraph data={exploitData} />
        <TotalPieGraph data={exploitData}/>
      </Box>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Tabs
          value={tabIndex}
          onChange={handleTabChange}
          aria-label="Graph View Tabs"
          variant="scrollable"
          scrollButtons="auto"
        >
          <Tab
            label="Rounds View"
            icon={<AccessTimeIcon />}
            iconPosition="start"
          />
          <Tab label="Teams View" icon={<GroupsIcon />} iconPosition="start" />
        </Tabs>
      </Box>

      <TabPanel value={tabIndex} index={0}>
        <TimelineGraph data={timelineData} />
      </TabPanel>

      <TabPanel value={tabIndex} index={1}>
        <AllTeamsGraph data={teamsData} />
      </TabPanel>
    </Box>
  );
}

export default App;

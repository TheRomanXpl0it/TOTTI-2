import { Box, Button, TextField } from "@mui/material";
import FlagIcon from '@mui/icons-material/Flag';
import { useState } from "react";

export default function ManualSubmission() {
  const [flag, setFlag] = useState("");
  const [highlight, setHighlight] = useState(false);

  const handleSend = () => {
    if (!flag.trim()) return;

    fetch("/api/manual", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify([{ flag }]),
    });

    setFlag("");
    setHighlight(true);
    setTimeout(() => setHighlight(false), 600);
  };

  const handleKeyPress = (event) => {
    if (event.key === "Enter") {
      handleSend();
    }
  };

  return (
    <Box display="flex" alignItems="center" width={400} gap={1}>
      <TextField
        size="small"
        label="Manual Flag"
        variant="outlined"
        fullWidth
        value={flag}
        onChange={(e) => setFlag(e.target.value)}
        onKeyPress={handleKeyPress}
        InputProps={{
          endAdornment: <FlagIcon fontSize="small" />,
        }}
        sx={{
          backgroundColor: highlight ? "rgb(71, 255, 121, 0.5)" : "inherit",
          transition: "background-color 0.3s ease",
        }}
      />
      <Button variant="contained" color="primary" onClick={handleSend}>
        SEND
      </Button>
    </Box>
  );
}

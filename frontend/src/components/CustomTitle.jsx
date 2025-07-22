import { Box, Typography } from "@mui/material";

function CustomTitle() {
  return (
    <Box>
      <Typography variant="h3" align="center" margin={0} padding={0}>
        <Box
          component="span"
          sx={{
            color: "primary.main",
            fontWeight: 700,
            textShadow: "0 0 15px rgb(94, 0, 0)",
          }}
          display={"flex"}
          flexDirection={"row"}
          justifyContent={"center"}
          alignItems={"center"}>
          <img
            src="totti.png"
            alt="Totti Logo"
            width="80"
            height="80"
            style={{ display: "block", margin: "0 auto" }}
          />
          TOTTI
        </Box>
      </Typography>
    </Box>
  );
}

export default CustomTitle;

import { createTheme } from '@mui/material/styles';

const darkRedTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      main: '#ffa200', // dark red
    },
    background: {
      default: '#121212',
      paper: '#1e1e1e',
    },
  },
});

export default darkRedTheme;

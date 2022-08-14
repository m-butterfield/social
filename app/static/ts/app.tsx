import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import {Header} from "components/header";
import React from "react";
import CssBaseline from "@mui/material/CssBaseline";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {theme} from "theme";
import {User} from "types";

declare const user: User | null;

const Home = React.lazy(() => import("components/home"));
const Login = React.lazy(() => import("components/login"));

export const App = () => {
  return <>
    <ThemeProvider theme={theme} >
      <CssBaseline />
      <BrowserRouter>
        <Header user={user} />
        <Container>
          <Box sx={{my: 2}}>
            <Routes>
              <Route
                path="/"
                element={
                  <React.Suspense fallback={<>...</>}>
                    <Home user={user} />
                  </React.Suspense>
                }
              />
              <Route
                path="login"
                element={
                  <React.Suspense fallback={<>...</>}>
                    <Login user={user} />
                  </React.Suspense>
                }
              />
              <Route path="*" element={<>not found...</>} />
            </Routes>
          </Box>
        </Container>
      </BrowserRouter>
    </ThemeProvider>
  </>;
};

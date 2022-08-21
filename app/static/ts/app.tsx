import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import {Header} from "components/header";
import {User} from "graphql/types";
import React, {createContext} from "react";
import CssBaseline from "@mui/material/CssBaseline";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {theme} from "theme";

declare const user: User | null;

export type AppState = {
  user: User | null;
}

export const AppContext = createContext<AppState>(null);

const Home = React.lazy(() => import("components/home"));
const Login = React.lazy(() => import("components/login"));

const client = new ApolloClient({
  uri: "http://localhost:8000/graphql/query",
  cache: new InMemoryCache(),
});

export const App = () => {
  return <AppContext.Provider value={{user: user}}>
    <ApolloProvider client={client}>
      <ThemeProvider theme={theme} >
        <CssBaseline />
        <BrowserRouter>
          <Header />
          <Container>
            <Box sx={{my: 2}}>
              <Routes>
                <Route
                  path="/"
                  element={
                    <React.Suspense fallback={<>...</>}>
                      <Home />
                    </React.Suspense>
                  }
                />
                <Route
                  path="login"
                  element={
                    <React.Suspense fallback={<>...</>}>
                      <Login />
                    </React.Suspense>
                  }
                />
                <Route path="*" element={<>not found...</>} />
              </Routes>
            </Box>
          </Container>
        </BrowserRouter>
      </ThemeProvider>
    </ApolloProvider>
  </AppContext.Provider>;
};

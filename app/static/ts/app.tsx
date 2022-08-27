import {ApolloClient, ApolloProvider, gql, InMemoryCache, useQuery} from "@apollo/client";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import Typography from "@mui/material/Typography";
import {Header} from "components/header";
import {Query, User} from "graphql/types";
import React, {createContext, useContext, useEffect, useState} from "react";
import CssBaseline from "@mui/material/CssBaseline";
import {BrowserRouter, Route, Routes} from "react-router-dom";
import {theme} from "theme";

export type AppState = {
  user: User | null;
  setUser: (user: User) => void;
}

export const AppContext = createContext<AppState>(null);

const Home = React.lazy(() => import("components/home"));
const Login = React.lazy(() => import("components/login"));

const client = new ApolloClient({
  uri: "http://localhost:8000/graphql/query",
  cache: new InMemoryCache(),
  credentials: "same-origin",
});

export const App = () => {
  const [stateUser, setUser] = useState<User | null>(null);
  const state = {
    user: stateUser,
    setUser: setUser,
  };
  return <AppContext.Provider value={state}>
    <ApolloProvider client={client}>
      <ThemeProvider theme={theme} >
        <CssBaseline />
        <Main />
      </ThemeProvider>
    </ApolloProvider>
  </AppContext.Provider>;
};


const ME = gql`
  query me {
    me {
      username
    }
  }
`;

const Main = () => {
  const {setUser} = useContext(AppContext);
  const {data, loading, error} = useQuery<Query>(ME);
  useEffect(() => {
    if (data && data.me) setUser(data.me);
  }, [data]);
  return <>
    {loading || error ?
      <Typography>{error && "something is wrong"}</Typography>
      :
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
    }
  </>;
};

import {ApolloClient, ApolloProvider, gql, InMemoryCache, useQuery} from "@apollo/client";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import Typography from "@mui/material/Typography";
import AppRoutes from "app/AppRoutes";
import {Header} from "app/Header";
import {Query, User} from "graphql/types";
import React, {createContext, useContext, useEffect, useState} from "react";
import CssBaseline from "@mui/material/CssBaseline";
import {createRoot} from "react-dom/client";
import {BrowserRouter} from "react-router-dom";
import {theme} from "app/theme";

type AppState = {
  user: User | null;
  setUser: (user: User) => void;
}

export const AppContext = createContext<AppState>(null);

const client = new ApolloClient({
  uri: "http://localhost:8000/graphql/query",
  cache: new InMemoryCache(),
  credentials: "same-origin",
});

const App = () => {
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
            <AppRoutes />
          </Box>
        </Container>
      </BrowserRouter>
    }
  </>;
};

const container = document.getElementById("root");
const root = createRoot(container);

root.render(<App />);

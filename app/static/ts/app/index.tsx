import {ApolloClient, ApolloProvider, InMemoryCache} from "@apollo/client";
import ThemeProvider from "@mui/material/styles/ThemeProvider";
import Main from "app/Main";
import {User} from "graphql/types";
import React, {createContext, useState} from "react";
import CssBaseline from "@mui/material/CssBaseline";
import {createRoot} from "react-dom/client";
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


const container = document.getElementById("root");
const root = createRoot(container);

root.render(<App />);

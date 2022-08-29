import {gql, useQuery} from "@apollo/client";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import AppRoutes from "app/AppRoutes";
import {Header} from "app/Header";
import {AppContext} from "app/index";
import {Query} from "graphql/types";
import React, {useContext, useEffect} from "react";
import {BrowserRouter} from "react-router-dom";

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

export default Main;

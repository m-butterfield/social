import {gql, useMutation} from "@apollo/client";
import AppBar from "@mui/material/AppBar";
import Container from "@mui/material/Container";
import Link from "@mui/material/Link";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import {AppContext} from "app";
import {Mutation} from "graphql/types";
import React, {useContext, useEffect, useState} from "react";
import {Link as RouterLink} from "react-router-dom";

const LOGOUT = gql`
  mutation logout {
    logout
  }
`;

export const Header = () => {
  const {user, setUser} = useContext(AppContext);
  const [logout] = useMutation<Mutation>(LOGOUT);
  const [logoutSuccess, setLogoutSuccess] = useState(false);

  useEffect(() => {
    if (!user && logoutSuccess) {
      window.location.href = "/";
    }
  }, [logoutSuccess]);

  return <AppBar
    position="static"
    color="secondary"
    enableColorOnDark={true}
    sx={{backgroundImage: "unset", boxShadow: "unset"}}
  >
    <Container>
      <Toolbar>
        <Typography variant="h6" sx={{flexGrow: 1}}>
          <Link underline="hover" color="text.primary" href="/">social</Link>
        </Typography>
        <nav>
          {user ?
            <>
              <Link
                component={RouterLink}
                underline="hover"
                color="text.primary"
                to="/create_post"
                sx={{my: 1, mx: 1.5}}
              >
                create post
              </Link>
              <Link
                underline="hover"
                color="text.primary"
                onClick={() => logout().then(() => {
                  setUser(null);
                  setLogoutSuccess(true);
                })}
                sx={{my: 1, mx: 1.5}}
              >
              logout
              </Link>
            </>
            :
            <Link
              component={RouterLink}
              underline="hover"
              color="text.primary"
              to="/login"
              sx={{my: 1, mx: 1.5}}
            >
            login / signup
            </Link>
          }
        </nav>
      </Toolbar>
    </Container>
  </AppBar>;
};

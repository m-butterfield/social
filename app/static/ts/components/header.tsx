import AppBar from "@mui/material/AppBar";
import Container from "@mui/material/Container";
import Link from "@mui/material/Link";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import {AppContext} from "app";
import React, {useContext} from "react";
import {Link as RouterLink} from "react-router-dom";

export const Header = () => {
  const {user} = useContext(AppContext);
  return <AppBar
    position="static"
    color="primary"
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
            <Link
              component={RouterLink}
              underline="hover"
              color="text.primary"
              to="/logout"
              sx={{my: 1, mx: 1.5}}
            >
            logout
            </Link>
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

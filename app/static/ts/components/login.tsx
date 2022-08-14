import Button from "@mui/material/Button";
import React from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import {User} from "types";

type LoginProps = {
  user?: User;
}

const Login = (props: LoginProps) => {
  const {user} = props;
  if (user) return <>"{user.id} You're already logged in!"</>;

  return <>
    <Typography align="center" variant="h3">login</Typography>
    <Box component="form" sx={{mt: 3}}>
      <Grid container spacing={2} alignItems="center" justifyContent="center" direction="column">
        <Grid item width={300}>
          <TextField
            label="username"
            fullWidth
          />
        </Grid>
        <Grid item width={300}>
          <TextField
            label="password"
            type="password"
            fullWidth
          />
        </Grid>
        <Grid item width={300}>
          <Button
            fullWidth type="submit"
            variant="contained"
            disabled={false}
            sx={{boxShadow: "unset"}}
          >
            login
          </Button>
        </Grid>
      </Grid>
    </Box>
    <Typography align="center" pt={4} variant="h3">or...</Typography>
    <Typography align="center" variant="h3">create account</Typography>
    <Box component="form" sx={{mt: 3}}>
      <Grid container spacing={2} alignItems="center" justifyContent="center" direction="column">
        <Grid item width={300}>
          <TextField
            label="username"
            fullWidth
          />
        </Grid>
        <Grid item width={300}>
          <TextField
            label="password"
            type="password"
            autoComplete="new-password"
            fullWidth
          />
        </Grid>
        <Grid item width={300}>
          <Button type="submit" disabled={false}>login</Button>
        </Grid>
      </Grid>
    </Box>
  </>;
};

export default Login;

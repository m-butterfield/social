import {gql, useMutation} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Button from "@mui/material/Button";
import {AppContext} from "app";
import Signup from "components/signup";
import {Mutation, MutationLoginArgs} from "graphql/types";
import React, {useContext, useState} from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import {Navigate} from "react-router-dom";

const LOGIN = gql`
  mutation login($input: UserCreds!) {
    login(input: $input) {
      username
    }
  }
`;

const Login = () => {
  const {user, setUser} = useContext(AppContext);
  const [success, setSuccess] = useState(false);

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const [login, {error}] = useMutation<
    Mutation, MutationLoginArgs
  >(LOGIN, {
    variables: {input: {username: username, password: password}}
  });

  if (user && !success) {
    return <>"{user.username} You're already logged in!"</>;
  }
  if (user && success) {
    return <Navigate to="/" />;
  }

  return <>
    <Typography align="center" variant="h4">login</Typography>
    <Box component="form" sx={{mt: 3}}>
      <Grid container spacing={2} alignItems="center" justifyContent="center" direction="column">
        <Grid item width={300}>
          <TextField
            fullWidth
            label="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
        </Grid>
        <Grid item width={300}>
          <TextField
            fullWidth
            label="password"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </Grid>
        {error &&
          <Grid item width={300}>
            <Alert severity="error">{error.message}</Alert>
          </Grid>
        }
        <Grid item width={300}>
          <Button
            fullWidth
            type="submit"
            variant="contained"
            disabled={false}
            sx={{boxShadow: "unset"}}
            onClick={(e) => {
              e.preventDefault();
              login().then((response) => {
                setUser(response.data.login);
                setSuccess(true);
              });
            }}
          >
            login
          </Button>
        </Grid>
      </Grid>
    </Box>
    <Typography align="center" pt={4} variant="h4">or...</Typography>
    <Signup setSuccess={setSuccess}/>
  </>;
};

export default Login;

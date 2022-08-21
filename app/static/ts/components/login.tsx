import {gql, useMutation} from "@apollo/client";
import Button from "@mui/material/Button";
import {AppContext} from "app";
import {Mutation, MutationCreateUserArgs} from "graphql/types";
import React, {useContext, useState} from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

const CREATE_USER = gql`
  mutation createUser($input: CreateUser!) {
    createUser(input: $input) {
      username
    }
  }
`;

const Login = () => {
  const {user} = useContext(AppContext);
  if (user) return <>"{user.username} You're already logged in!"</>;

  const [newUsername, setNewUsername] = useState("");
  const [newPassword, setNewPassword] = useState("");

  const [createUser, {error, data}] = useMutation<
    Mutation, MutationCreateUserArgs
  >(CREATE_USER, {
    variables: {input: {username: newUsername, password: newPassword}}
  });

  return <>
    <Typography align="center" variant="h4">login</Typography>
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
    <Typography align="center" pt={4} variant="h4">or...</Typography>
    <Typography align="center" variant="h4">create account</Typography>
    <Box component="form" sx={{mt: 3}}>
      <Grid container spacing={2} alignItems="center" justifyContent="center" direction="column">
        <Grid item width={300}>
          <TextField
            label="username"
            onChange={(e) => setNewUsername(e.target.value)}
            fullWidth
          />
        </Grid>
        <Grid item width={300}>
          <TextField
            label="password"
            type="password"
            autoComplete="new-password"
            onChange={(e) => setNewPassword(e.target.value)}
            fullWidth
          />
        </Grid>
        <Grid item width={300}>
          <Button
            fullWidth type="submit"
            variant="contained"
            disabled={false}
            sx={{boxShadow: "unset"}}
            onClick={(e) => {
              e.preventDefault();
              createUser().catch();
            }}
          >
            create account
          </Button>
        </Grid>
      </Grid>
    </Box>
  </>;
};

export default Login;

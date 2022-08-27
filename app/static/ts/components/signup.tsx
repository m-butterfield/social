import {gql, useMutation} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Button from "@mui/material/Button";
import {AppContext} from "app";
import {Mutation, MutationCreateUserArgs} from "graphql/types";
import React, {useContext, useState} from "react";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";

const CREATE_USER = gql`
  mutation createUser($input: UserCreds!) {
    createUser(input: $input) {
      username
    }
  }
`;

type SignupProps = {
  setSuccess: (s: boolean) => void;
};

const Signup = (props: SignupProps) => {
  const {setUser} = useContext(AppContext);
  const [newUsername, setNewUsername] = useState("");
  const [newPassword, setNewPassword] = useState("");

  const [createUser, {error}] = useMutation<
    Mutation, MutationCreateUserArgs
  >(CREATE_USER, {
    variables: {input: {username: newUsername, password: newPassword}}
  });

  return <>
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
        {error &&
          <Grid item width={300}>
            <Alert severity="error">{error.message}</Alert>
          </Grid>
        }
        <Grid item width={300}>
          <Button
            fullWidth type="submit"
            variant="contained"
            disabled={false}
            sx={{boxShadow: "unset"}}
            onClick={(e) => {
              e.preventDefault();
              createUser().then((response) => {
                props.setSuccess(true);
                setUser(response.data.createUser);
              });
            }}
          >
            create account
          </Button>
        </Grid>
      </Grid>
    </Box>
  </>;
};

export default Signup;

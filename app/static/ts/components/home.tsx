import React from "react";
import Typography from "@mui/material/Typography";
import {User} from "types";

type HomeProps = {
  user?: User;
}

const Home = (props: HomeProps) => {
  const {user} = props;
  return <Typography variant="h2">Welcome {user && user.id}</Typography>;
};

export default Home;

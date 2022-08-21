import {AppContext} from "app";
import React, {useContext} from "react";
import Typography from "@mui/material/Typography";

const Home = () => {
  const {user} = useContext(AppContext);
  return <Typography variant="h2">Welcome {user && user.username}</Typography>;
};

export default Home;

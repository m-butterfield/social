import {useQuery} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Typography from "@mui/material/Typography";
import {AppContext} from "app/index";
import ScrollablePosts from "app/lib/components/ScrollablePosts";
import {FollowButton} from "app/user_home/FollowButton";
import {GET_USER, GET_USER_POSTS} from "app/user_home/queries";
import React, {useContext} from "react";
import {useParams} from "react-router-dom";


const UserHome = () => {
  const {user} = useContext(AppContext);
  const {username} = useParams();

  const {data, loading, error} = useQuery(GET_USER, {variables: {username: username}});

  if (loading) {
    return <Typography>Loading...</Typography>;
  }
  if (error) {
    return <Alert severity="error">{error.message}</Alert>;
  }

  return <ScrollablePosts
    query={GET_USER_POSTS}
    queryName={"getUserPosts"}
    queryVariables={{username: username}}
    header={<>
      <Typography variant="h2">{username}</Typography>
      {user && user.id !== data.getUser.id &&
       <FollowButton homeUser={data.getUser} />
      }
    </>}
  />;
};

export default UserHome;

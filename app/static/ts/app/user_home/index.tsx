import {useQuery} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import {AppContext} from "app/index";
import PostItem from "app/lib/components/PostItem";
import {FollowButton} from "app/user_home/FollowButton";
import {GET_USER_POSTS} from "app/user_home/queries";
import {Post} from "graphql/types";
import React, {useContext} from "react";
import {useParams} from "react-router-dom";


const UserHome = () => {
  const {user} = useContext(AppContext);
  const {username} = useParams();
  const {data, loading, error} = useQuery(GET_USER_POSTS, {variables: {username: username}});

  if (loading) {
    return <Typography>Loading...</Typography>;
  }
  if (error) {
    return <Alert severity="error">Error loading posts</Alert>;
  }

  return <Stack direction="column" alignItems="center" spacing={2} width={800} m="auto">
    <Typography variant="h2">{username}</Typography>
    {user.id !== data.getUserPosts.user.id &&
      <FollowButton
        homeUser={data.getUserPosts.user}
      />
    }
    {data.getUserPosts.posts.map((post: Post) => {
      return <PostItem key={post.id} post={post} />;
    })}
  </Stack>;
};

export default UserHome;

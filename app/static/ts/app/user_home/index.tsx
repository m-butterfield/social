import {gql, useMutation, useQuery} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import {AppContext} from "app/index";
import PostItem from "app/lib/components/PostItem";
import {Mutation, Post, MutationFollowUserArgs} from "graphql/types";
import React, {useContext} from "react";
import {useParams} from "react-router-dom";

const GET_USER_POSTS = gql`
  query getUserPosts($username: String!) {
    getUserPosts(username: $username) {
      posts {
        id
        body
        images {
          id
          width
          height
        }
        user {
          username
        }
      }
      user {
        id
      }
    }
  }
`;

const FOLLOW_USER = gql`
  mutation followUser($username: String!) {
    followUser(username: $username)
  }
`;

const UserHome = () => {
  const {user} = useContext(AppContext);
  const {username} = useParams();

  const {data, loading, error} = useQuery(GET_USER_POSTS, {variables: {username: username}});
  const message = loading ? "Loading..." : error ? "Error loading posts..." : "";

  const [followUser, {error: followError}] = useMutation<
    Mutation, MutationFollowUserArgs
  >(FOLLOW_USER, {
    variables: {username: username}
  });

  const isFollowingUser = data && user.following.find(f => f.userID === data.getUserPosts.user.id);

  if (loading) {
    return <Typography>Loading...</Typography>;
  }

  return <Stack direction="column" alignItems="center" spacing={2} width={800} m="auto">
    <Typography variant="h2">{username}</Typography>

    {followError && <Alert severity="error">Error following user: {followError.message}</Alert>}
    {isFollowingUser ?
      <Button>Unfollow</Button>
      :
      <Button
        type="submit"
        variant="contained"
        onClick={() => followUser()}
      >
      Follow
      </Button>
    }

    {message ?
      <Typography>{message}</Typography>
      :
      data.getUserPosts.posts.map((post: Post) => {
        return <PostItem key={post.id} post={post} />;
      })
    }
  </Stack>;
};

export default UserHome;

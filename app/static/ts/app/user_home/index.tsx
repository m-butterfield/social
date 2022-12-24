import {gql, useQuery} from "@apollo/client";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import PostItem from "app/lib/components/PostItem";
import {Post} from "graphql/types";
import React from "react";
import {useParams} from "react-router-dom";

const GET_POSTS = gql`
  query getUserPosts($userID: String!) {
    getUserPosts(userID: $userID) {
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
  }
`;

const UserHome = () => {
  const {userID} = useParams();

  const {data, loading, error} = useQuery(GET_POSTS, {variables: {userID: userID}});
  const message = loading ? "Loading..." : error ? "Error loading posts..." : "";

  return <Stack direction="column" alignItems="center" spacing={2} width={800} m="auto">
    <Typography variant="h2">{userID}</Typography>
    {message ?
      <Typography>{message}</Typography>
      :
      data.getUserPosts.map((post: Post) => {
        return <PostItem key={post.id} post={post} />;
      })
    }
  </Stack>;
};

export default UserHome;

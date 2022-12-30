import {gql, useQuery} from "@apollo/client";
import Stack from "@mui/material/Stack";
import {AppContext} from "app";
import PostItem from "app/lib/components/PostItem";
import {Post} from "graphql/types";
import React, {useContext} from "react";
import Typography from "@mui/material/Typography";

const GET_NEW_POSTS = gql`
  query getNewPosts {
    getNewPosts {
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

const New = () => {
  const {user} = useContext(AppContext);

  const {data, loading, error} = useQuery(GET_NEW_POSTS);
  const message = loading ? "Loading..." : error ? "Error loading posts..." : "";

  return <Stack direction="column" alignItems="center" spacing={2} width={800} m="auto">
    <Typography variant="h2">Welcome{user && `, ${user.username}.`}</Typography>
    {message ?
      <Typography>{message}</Typography>
      :
      data.getNewPosts.map((post: Post) => {
        return <PostItem key={post.id} post={post} />;
      })
    }
  </Stack>;
};

export default New;

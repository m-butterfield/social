import {gql, useQuery} from "@apollo/client";
import Stack from "@mui/material/Stack";
import {debounce} from "@mui/material/utils";
import {AppContext} from "app";
import PostItem from "app/lib/components/PostItem";
import {Post} from "graphql/types";
import React, {useContext, useEffect, useState} from "react";
import Typography from "@mui/material/Typography";

const GET_NEW_POSTS = gql`
  query getNewPosts($before: Time) {
    getNewPosts(before: $before) {
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
      publishedAt
    }
  }
`;

const New = () => {
  const {user} = useContext(AppContext);

  const {data, loading, error, refetch} = useQuery(GET_NEW_POSTS);
  const message = loading ? "Loading..." : error ? "Error loading posts..." : "";
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    if (data?.getNewPosts) {
      setPosts([...posts, ...data.getNewPosts]);

      window.addEventListener("scroll", debounce(async () => {
        if ((window.innerHeight + window.scrollY) >= document.body.offsetHeight - 50) {
          const before = data.getNewPosts[data.getNewPosts.length - 1].publishedAt;
          await refetch({before: before});
        }
      }, 250));
    }
  }, [data]);

  return <Stack direction="column" alignItems="center" spacing={2} width={800} m="auto">
    <Typography variant="h2">Welcome{user && `, ${user.username}.`}</Typography>
    {message ?
      <Typography>{message}</Typography>
      :
      posts.map((post: Post) => {
        return <PostItem key={post.id} post={post} />;
      })
    }
  </Stack>;
};

export default New;

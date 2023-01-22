import {gql, useQuery} from "@apollo/client";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import {setRef} from "@mui/material/utils";
import {AppContext} from "app";
import PostItem from "app/lib/components/PostItem";
import {Post} from "graphql/types";
import React, {useCallback, useContext, useEffect, useRef, useState} from "react";

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
  const [noNewData, setNoNewData] = useState(false);
  const [refetched, setRefetched] = useState(false);

  useEffect(() => {
    if (data?.getNewPosts) {
      if (!data.getNewPosts.length) {
        setNoNewData(true);
      } else {
        setPosts([...posts, ...data.getNewPosts]);
      }
    }
  }, [data]);

  useEffect(() => {
    if (noNewData) {
      return;
    }
    const refetchEvent = async () => {
      if (!refetched && window.innerHeight + window.scrollY >= document.body.offsetHeight - 50) {
        setRefetched(true);
        const before = posts[posts.length - 1].publishedAt;
        await refetch({before: before});
        setRefetched(false);
      }
    };
    window.addEventListener("scroll", refetchEvent);
    return () => window.removeEventListener("scroll", refetchEvent);
  }, [noNewData, posts]);

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

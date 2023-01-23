import {gql, NetworkStatus, useQuery} from "@apollo/client";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import {AppContext} from "app";
import PostItem from "app/lib/components/PostItem";
import {Post} from "graphql/types";
import React, {useContext, useEffect, useState} from "react";

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

  const {data, loading, error, refetch, networkStatus} = useQuery(
    GET_NEW_POSTS, {notifyOnNetworkStatusChange: true}
  );
  const refetching = networkStatus === NetworkStatus.setVariables; // this should be NetworkStatus.refetch but there is a bug: https://github.com/apollographql/apollo-client/issues/10391
  const message = loading && !refetching ? "Loading..." : error ? "Error loading posts..." : "";
  const [posts, setPosts] = useState<Post[]>([]);
  const [noNewData, setNoNewData] = useState(false);
  const [before, setBefore] = useState("");

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
    if (before) {
      refetch({before: before});
    }
  }, [before]);

  useEffect(() => {
    if (noNewData) {
      return;
    }
    const refetchEvent = async () => {
      if (window.innerHeight + window.scrollY >= document.body.offsetHeight - 200) {
        setBefore(posts[posts.length - 1].publishedAt);
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
    {refetching && <Typography>Loading more posts...</Typography>}
    {noNewData && <Typography>No more posts.</Typography>}
  </Stack>;
};

export default New;

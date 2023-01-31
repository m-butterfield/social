import {gql} from "@apollo/client";
import ScrollablePosts from "app/lib/components/ScrollablePosts";
import React from "react";

const GET_POSTS = gql`
  query getPosts($before: Time) {
    getPosts(before: $before) {
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
      commentCount
      film
      camera
      lens
    }
  }
`;

const Home = () => {
  return <ScrollablePosts
    query={GET_POSTS}
    queryName={"getPosts"}
  />;
};

export default Home;

import {gql} from "@apollo/client";
import ScrollablePosts from "app/lib/components/ScrollablePosts";
import React from "react";

const GET_POSTS = gql`
  query getPostsCoolStuff($before: Time) {
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

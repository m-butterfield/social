import {gql} from "@apollo/client";
import ScrollablePosts from "app/lib/components/ScrollablePosts";
import React from "react";

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
      commentCount
    }
  }
`;

const New = () => {
  return <ScrollablePosts
    query={GET_NEW_POSTS}
    queryName={"getNewPosts"}
  />;
};

export default New;

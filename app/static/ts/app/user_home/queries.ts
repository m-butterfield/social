import {gql} from "@apollo/client";

export const GET_USER_POSTS = gql`
  query getUserPosts($username: String!, $before: Time) {
    getUserPosts(username: $username, before: $before) {
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

export const GET_USER = gql`
  query getUser($username: String!) {
    getUser(username: $username) {
      id
      username
    }
  }
`;

export const FOLLOW_USER = gql`
  mutation followUser($username: String!) {
    followUser(username: $username)
  }
`;

export const UNFOLLOW_USER = gql`
  mutation unFollowUser($username: String!) {
    unFollowUser(username: $username)
  }
`;

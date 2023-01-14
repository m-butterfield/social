import {gql} from "@apollo/client";

export const GET_USER_POSTS = gql`
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
        username
      }
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

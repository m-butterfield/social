import {gql} from "@apollo/client";

export const GET_COMMENTS = gql`
  query getComments($postID: String!, $before: Time) {
    getComments(postID: $postID, before: $before) {
      id
      body
      user {
        username
      }
    }
  }
`;

export const CREATE_COMMENT = gql`
  mutation createComment($postID: String!, $body: String!) {
    createComment(postID: $postID, body: $body) {
      id
    }
  }
`;

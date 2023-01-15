import {gql} from "@apollo/client";

export const ME = gql`
  query me {
    me {
      id
      username
      following {
        userID
      }
    }
  }
`;

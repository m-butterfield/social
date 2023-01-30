import {useQuery} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Link from "@mui/material/Link";
import Typography from "@mui/material/Typography";
import {GET_COMMENTS} from "app/lib/components/PostItem/queries";
import {Comment, Post} from "graphql/types";
import React from "react";

type CommentsProps = {
  post: Post
}

const PostItem = (props: CommentsProps) => {
  const {post} = props;
  const {loading, error, data} = useQuery(GET_COMMENTS, {
    variables: {
      postID: post.id,
    }
  });
  if (error) {
    return  <Alert severity="error">Error loading comments: {error.message}</Alert>;
  }
  if (loading) {
    return <Typography>Loading comments...</Typography>;
  }
  return data.getComments.map((comment: Comment) =>
    <Typography>
      <Link href={`/${comment.user.username}`}>{comment.user.username}</Link>: {comment.body}
    </Typography>
  );
};

export default PostItem;

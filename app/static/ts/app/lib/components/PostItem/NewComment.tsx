import {useMutation} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import {CREATE_COMMENT} from "app/lib/components/PostItem/queries";
import {Mutation, MutationCreateCommentArgs, Post} from "graphql/types";
import React, {useState} from "react";

type CommentsProps = {
  post: Post;
  refetch: () => void;
}

const NewComment = (props: CommentsProps) => {
  const {post, refetch} = props;
  const [comment, setComment] = useState("");

  const [createComment, {loading, error}] = useMutation<
    Mutation, MutationCreateCommentArgs
  >(CREATE_COMMENT, {
    variables: {postID: post.id, body: comment}
  });

  return <>
    {error && <Alert severity="error">Error saving comment: {error.message}</Alert>}
    <TextField
      label="comment"
      multiline
      maxRows={20}
      fullWidth
      placeholder="write a new comment..."
      value={comment}
      onChange={(e) => setComment(e.target.value)}
    />
    <Button
      type="submit"
      variant="contained"
      disabled={!comment || loading}
      onClick={(e) => {
        e.preventDefault();
        createComment().then(() => {
          refetch();
          setComment("");
        });
      }}
    >
      submit
    </Button>
  </>;
};

export default NewComment;

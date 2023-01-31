import {gql, useMutation} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import Link from "@mui/material/Link";
import {Mutation, MutationDeletePostArgs, Post} from "graphql/types";
import React, {useState} from "react";

const DELETE_POST = gql`
  mutation deletePost($postID: String!) {
    deletePost(postID: $postID)
  }
`;

type DeletePostProps = {
  post: Post
  posts: Post[];
  setPosts: (v: Post[]) => void;
}

const DeletePostButton = (props: DeletePostProps) => {
  const {post, posts, setPosts} = props;
  const [modalOpen, setModalOpen] = useState(false);

  const [deletePost, {loading, error}] = useMutation<
    Mutation, MutationDeletePostArgs
  >(DELETE_POST, {
    variables: {postID: post.id}
  });

  if (error) {
    return <Alert severity="error">{error.message}</Alert>;
  }

  return <>
    <Link
      component="button"
      fontSize="1rem"
      onClick={() => setModalOpen(true)}
    >
      delete
    </Link>
    <Dialog
      open={modalOpen}
      onClose={() => setModalOpen(false)}
    >
      <DialogContent>
        <DialogContentText>
          Are you sure you want to delete this post?
        </DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button onClick={() => setModalOpen(false)}>Cancel</Button>
        <Button
          disabled={loading}
          onClick={() => {
            deletePost().then(() => {
              setModalOpen(false);
              setPosts(posts.filter((p) => p.id !== post.id));
            });
          }}>Delete Post</Button>
      </DialogActions>
    </Dialog>
  </>;
};

export default DeletePostButton;

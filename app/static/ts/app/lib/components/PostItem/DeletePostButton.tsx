import Button from "@mui/material/Button";
import Dialog from "@mui/material/Dialog";
import DialogActions from "@mui/material/DialogActions";
import DialogContent from "@mui/material/DialogContent";
import DialogContentText from "@mui/material/DialogContentText";
import Link from "@mui/material/Link";
import {AppContext} from "app/index";
import {Post} from "graphql/types";
import React, {useContext, useState} from "react";

type DeletePostProps = {
  post: Post
}

const DeletePostButton = (props: DeletePostProps) => {
  const {user} = useContext(AppContext);
  const {post} = props;
  const [modalOpen, setModalOpen] = useState(false);
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
        <Button onClick={() => {
          setModalOpen(false);
        }}>Delete Post</Button>
      </DialogActions>
    </Dialog>
  </>;
};

export default DeletePostButton;

import {useQuery} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Divider from "@mui/material/Divider";
import Link from "@mui/material/Link";
import Typography from "@mui/material/Typography";
import {AppContext} from "app/index";
import NewComment from "app/lib/components/PostItem/NewComment";
import {GET_COMMENTS} from "app/lib/components/PostItem/queries";
import {Comment, Post} from "graphql/types";
import React, {useContext} from "react";

type CommentsProps = {
  post: Post
  setNumComments: (val: number) => void
}

const Comments = (props: CommentsProps) => {
  const {post, setNumComments} = props;
  const {user} = useContext(AppContext);
  const {loading, error, data, refetch} = useQuery(GET_COMMENTS, {
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
  setNumComments(data.getComments.length);
  return <>
    {
      data.getComments.map((comment: Comment) =>
        <>
          <Divider />
          <Typography>
            <Link href={`/${comment.user.username}`}>{comment.user.username}</Link>: {comment.body}
          </Typography>
        </>
      )
    }
    {user ?
      <NewComment post={post} refetch={refetch} />
      :
      <Typography variant="subtitle2"><em>You must be <Link href="/login">logged</Link> in to comment.</em></Typography>
    }
  </>;
};

export default Comments;

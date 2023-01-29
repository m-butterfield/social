import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import {IMAGES_BASE_URL} from "app/lib/constants";
import {Post} from "graphql/types";
import React, {useState} from "react";

type PostProps = {
  post: Post
}

const PostItem = (props: PostProps) => {
  const {post} = props;
  const [showComments, setShowComments] = useState(true);
  const [comment, setComment] = useState("");
  return <Box sx={{paddingBottom: 10}}>
    <Card sx={{paddingX: 4, paddingY: 2, borderRadius: 0, backgroundColor: "#303030"}}>
      <Stack direction="column" alignItems="center" spacing={2} m="auto">
        <Typography><Link href={`/${post.user.username}`}>{post.user.username}</Link></Typography>
        {
          post.images.map((image) => {
            return <img
              key={image.id}
              src={`${IMAGES_BASE_URL}${image.id}`}
              alt="post image"
              style={{width: image.width, maxHeight: image.height}}
            />;
          })
        }
        <Typography>{post.body}</Typography>
        <Link
          component="button"
          fontSize="1rem"
          onClick={() => {
            if (showComments) {
              setShowComments(false);
            } else {
              setShowComments(true);
            }
          }}
        >0 comments</Link>
        {showComments &&
          <>
            <Typography><Link href={"#"}>commenter</Link>: This is a good post.</Typography>
            <TextField
              label="comment"
              multiline
              maxRows={20}
              sx={{width: "75ch"}}
              placeholder="write a new comment..."
              value={comment}
              onChange={(e) => setComment(e.target.value)}
            />
            <Button
              type="submit"
              variant="contained"
              disabled={!comment}
              onClick={(e) => {
                e.preventDefault();
                setComment("");
              }}
            >
              submit
            </Button>
          </>
        }
      </Stack>
    </Card>
  </Box>;
};

export default PostItem;

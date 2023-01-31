import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import {AppContext} from "app/index";
import Comments from "app/lib/components/PostItem/Comments";
import DeletePostButton from "app/lib/components/PostItem/DeletePostButton";
import {IMAGES_BASE_URL} from "app/lib/constants";
import {Post} from "graphql/types";
import React, {useContext, useState} from "react";

type PostProps = {
  post: Post
}

const PostItem = (props: PostProps) => {
  const {post} = props;
  const {user} = useContext(AppContext);
  const [showComments, setShowComments] = useState(false);
  const [numComments, setNumComments] = useState(post.commentCount);
  return <Box sx={{paddingBottom: 10}}>
    <Card sx={{paddingX: 4, paddingY: 4, borderRadius: 0, backgroundColor: "#303030"}}>
      <Stack direction="column" alignItems="flex-start" spacing={2}>
        <Stack direction="row" justifyContent="space-between" width="100%">
          <Link href={`/${post.user.username}`}>{post.user.username}</Link>
          {user && <DeletePostButton post={post} />}
        </Stack>
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
        {post.body && <Typography>{post.body}</Typography>}
        {post.film && <Typography variant="subtitle2">Film: {post.film}</Typography>}
        {post.camera && <Typography variant="subtitle2">Camera: {post.camera}</Typography>}
        {post.lens && <Typography variant="subtitle2">Lens: {post.lens}</Typography>}
        {showComments ?
          <>
            <Link
              component="button"
              fontSize="1rem"
              onClick={() => setShowComments(false)}
            >hide comments</Link>
            <Box sx={{width: "100%"}}>
              <Stack direction="column" spacing={2}>
                <Comments post={post} setNumComments={setNumComments} />
              </Stack>
            </Box>
          </>
          :
          <Link
            component="button"
            fontSize="1rem"
            onClick={() => setShowComments(true)}
          >{numComments ? `view ${numComments} comments` : "add comment"}</Link>
        }
      </Stack>
    </Card>
  </Box>;
};

export default PostItem;

import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import Comments from "app/lib/components/PostItem/Comments";
import {IMAGES_BASE_URL} from "app/lib/constants";
import {Post} from "graphql/types";
import React, {useState} from "react";

type PostProps = {
  post: Post
}

const PostItem = (props: PostProps) => {
  const {post} = props;
  const [showComments, setShowComments] = useState(false);
  const [numComments, setNumComments] = useState(post.commentCount);
  return <Box sx={{paddingBottom: 10}}>
    <Card sx={{paddingX: 4, paddingY: 4, borderRadius: 0, backgroundColor: "#303030"}}>
      <Stack direction="column" alignItems="flex-start" spacing={2}>
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
        <Typography><Link href={`/${post.user.username}`}>{post.user.username}</Link>: {post.body}</Typography>
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

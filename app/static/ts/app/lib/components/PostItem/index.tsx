import Box from "@mui/material/Box";
import Card from "@mui/material/Card";
import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import Comments from "app/lib/components/PostItem/Comments";
import NewComment from "app/lib/components/PostItem/NewComment";
import {IMAGES_BASE_URL} from "app/lib/constants";
import {Post} from "graphql/types";
import React, {useState} from "react";

type PostProps = {
  post: Post
}

const PostItem = (props: PostProps) => {
  const {post} = props;
  const [showComments, setShowComments] = useState(false);
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
        {showComments ?
          <>
            <Comments post={post} />
            <NewComment post={post} />
          </>
          :
          <Link
            component="button"
            fontSize="1rem"
            onClick={() => setShowComments(true)}
          >view comments</Link>
        }
      </Stack>
    </Card>
  </Box>;
};

export default PostItem;

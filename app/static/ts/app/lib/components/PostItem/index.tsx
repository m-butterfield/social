import Link from "@mui/material/Link";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import {IMAGES_BASE_URL} from "app/lib/constants";
import {Post} from "graphql/types";
import React from "react";

type PostProps = {
  post: Post
}

const PostItem = (props: PostProps) => {
  const {post} = props;
  return <Stack direction="column" alignItems="center" spacing={2} m="auto" sx={{paddingBottom: 20}}>
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
  </Stack>;
};

export default PostItem;

import Typography from "@mui/material/Typography";
import {IMAGES_BASE_URL} from "app/lib/constants";
import {Post} from "graphql/types";
import React from "react";

type PostProps = {
  post: Post
}

const PostItem = (props: PostProps) => {
  const {post} = props;
  return <>
    <Typography>{post.user.username}</Typography>
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
  </>;
};

export default PostItem;

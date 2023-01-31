import {gql, useMutation} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";
import Stack from "@mui/material/Stack";
import Typography from "@mui/material/Typography";
import PostTextField from "app/create_post/PostTextField";
import {CreatePostInput, Mutation, MutationCreatePostArgs, MutationSignedUploadUrlArgs} from "graphql/types";
import React, {useEffect, useState} from "react";

const MAX_POST_LENGTH = 4096;
const MAX_PROP_LENGTH = 255;

const CREATE_POST = gql`
  mutation createPost($input: CreatePostInput!) {
    createPost(input: $input) {
      id
    }
  }
`;

const SIGNED_UPLOAD_URL = gql`
  mutation signedUploadURL($input: SignedUploadInput!) {
    signedUploadURL(input: $input)
  }
`;

const uploadFile = (url: string, imageFile: File) => {
  return fetch(url, {
    method: "PUT",
    headers: new Headers({"Content-Type": imageFile.type}),
    body: imageFile,
  });
};

const INITIAL_INPUT: CreatePostInput = {
  body: "",
  film: "",
  camera: "",
  lens: "",
  images: [],
};

const CreatePost = () => {
  const [file, setFile] = useState(null);
  const [fileName, setFileName] = useState("");
  const [hasError, setHasError] = useState(false);
  const [postID, setPostID] = useState("");
  const [loading, setLoading] = useState(false);
  const [input, setInput] = useState<CreatePostInput>(INITIAL_INPUT);

  useEffect(() => {
    setInput({
      ...input,
      images: [fileName ? `${fileName}?${crypto.randomUUID()}` : ""],
    });
  }, [fileName]);

  const [createPost, {error: postError}] = useMutation<
    Mutation, MutationCreatePostArgs
  >(CREATE_POST, {
    variables: {input}
  });

  const [signedUploadURL, {error: uploadError}] = useMutation<
    Mutation, MutationSignedUploadUrlArgs
  >(SIGNED_UPLOAD_URL, {
    variables: {
      input: {
        fileName: input?.images.length ? input.images[0] : "",
        contentType: file ? file.type : "",
      }
    }
  });

  const error = postError || uploadError;

  return <>
    <Typography align="center" variant="h4">create a post</Typography>
    <Box component="form" sx={{mt: 3}}>
      <Stack direction="column" alignItems="center" spacing={2} width={300} m="auto">
        <Button fullWidth variant="contained" component="label" disabled={loading || postID !== ""}>
          Choose Image File
          <input hidden accept="image/jpeg" type="file" onChange={(e) => {
            if (e.target.files.length) {
              setFileName(e.target.files[0].name);
              setFile(e.target.files[0]);
            } else {
              setFileName("");
              setFile(null);
            }
          }}/>
        </Button>
        {fileName &&
          <Typography>File chosen: {fileName}</Typography>
        }
        <PostTextField
          disabled={loading || postID !== ""}
          maxLength={MAX_POST_LENGTH}
          input={input}
          setInput={setInput}
          field="body"
          label="caption"
          setHasError={setHasError}
        />
        <PostTextField
          disabled={loading || postID !== ""}
          maxLength={MAX_PROP_LENGTH}
          input={input}
          setInput={setInput}
          field="film"
          label="film type"
          setHasError={setHasError}
        />
        <PostTextField
          disabled={loading || postID !== ""}
          maxLength={MAX_PROP_LENGTH}
          input={input}
          setInput={setInput}
          field="camera"
          label="camera"
          setHasError={setHasError}
        />
        <PostTextField
          disabled={loading || postID !== ""}
          maxLength={MAX_PROP_LENGTH}
          input={input}
          setInput={setInput}
          field="lens"
          label="lens"
          setHasError={setHasError}
        />
        {error &&
          <Alert severity="error">{error.message}</Alert>
        }
        {postID &&
          <Alert severity="success">Success!</Alert>
        }
        <Button
          fullWidth
          variant="contained"
          disabled={!fileName || hasError || loading || postID !== ""}
          component="label"
          onClick={(e) => {
            e.preventDefault();
            setLoading(true);
            signedUploadURL()
              .then((r) => uploadFile(r.data.signedUploadURL, file))
              .then(() => createPost())
              .then((response) => setPostID(response.data.createPost.id))
              .finally(() => setLoading(false));
          }}
        >
          Submit
        </Button>
      </Stack>
    </Box>
  </>;
};

export default CreatePost;

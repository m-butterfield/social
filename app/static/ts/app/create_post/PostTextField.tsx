import TextField from "@mui/material/TextField";
import {CreatePostInput} from "graphql/types";
import React, {useState} from "react";

type PostTextFieldProps = {
  disabled: boolean;
  maxLength: number;
  input: CreatePostInput;
  setInput: (v: CreatePostInput) => void;
  label: string;
  field: "body" | "film" | "camera" | "lens";
  setHasError: (v: boolean) => void;
}

const PostTextField = (props: PostTextFieldProps) => {
  const {
    disabled,
    maxLength,
    input,
    setInput,
    label,
    field,
    setHasError,
  } = props;
  const [error, setError] = useState("");

  return <TextField
    multiline
    fullWidth
    disabled={disabled}
    label={label}
    error={error.length > 0}
    helperText={error}
    value={input && input[field]}
    onChange={(e) => {
      if (e.target.value.length > maxLength) {
        setError("too long, max 4096 characters");
        setHasError(true);
      } else if (error.length) {
        setError("");
        setHasError(false);
      }
      setInput({
        ...input,
        [field]: e.target.value,
      });
    }}
  ></TextField>;
};

export default PostTextField;

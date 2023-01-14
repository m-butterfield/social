import {useMutation} from "@apollo/client";
import Alert from "@mui/material/Alert";
import Button from "@mui/material/Button";
import {AppContext} from "app/index";
import {FOLLOW_USER, UNFOLLOW_USER} from "app/user_home/queries";
import {Mutation, MutationFollowUserArgs, MutationUnFollowUserArgs, User} from "graphql/types";
import React, {useContext} from "react";

type FollowButtonProps = {
  homeUser: User
}

export const FollowButton = ({homeUser}: FollowButtonProps) => {
  const {user, setUser} = useContext(AppContext);

  const [followUser, {error: followError}] = useMutation<
    Mutation, MutationFollowUserArgs
  >(FOLLOW_USER, {
    variables: {username: homeUser.username}
  });

  const [unFollowUser, {error: unFollowError}] = useMutation<
    Mutation, MutationUnFollowUserArgs
  >(UNFOLLOW_USER, {
    variables: {username: homeUser.username}
  });

  const isFollowingUser = homeUser && user.following.find(f => f.userID === homeUser.id);

  const updateUser = (isFollowing: boolean) => {
    if (isFollowing && !isFollowingUser) {
      setUser({
        ...user,
        following: [
          ...user.following,
          {userID: homeUser.id},
        ]
      });
    } else if (!isFollowing && isFollowingUser) {
      setUser({
        ...user,
        following: [
          ...user.following.filter((f) => f.userID !== homeUser.id),
        ]
      });
    }
  };

  return <>
    {followError && <Alert severity="error">Error following user: {followError.message}</Alert>}
    {unFollowError && <Alert severity="error">Error unfollowing user: {unFollowError.message}</Alert>}
    {isFollowingUser ?
      <Button
        type="submit"
        variant="contained"
        onClick={() => unFollowUser().then(() => updateUser(false))}
      >
      Unfollow
      </Button>
      :
      <Button
        type="submit"
        variant="contained"
        onClick={() => followUser().then(() => updateUser(true))}
      >
      Follow
      </Button>
    }
  </>;
};

export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: any;
};

export type Comment = {
  __typename?: 'Comment';
  body: Scalars['String'];
  createdAt: Scalars['Time'];
  id: Scalars['String'];
  user: User;
};

export type CreatePostInput = {
  body: Scalars['String'];
  camera: Scalars['String'];
  film: Scalars['String'];
  images: Array<Scalars['String']>;
  lens: Scalars['String'];
};

export type Follow = {
  __typename?: 'Follow';
  userID: Scalars['String'];
};

export type GetUserInput = {
  username: Scalars['String'];
};

export type Image = {
  __typename?: 'Image';
  height: Scalars['Int'];
  id: Scalars['String'];
  width: Scalars['Int'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createComment: Comment;
  createPost: Post;
  createUser: User;
  followUser: Scalars['Boolean'];
  login: User;
  logout: Scalars['Boolean'];
  signedUploadURL: Scalars['String'];
  unFollowUser: Scalars['Boolean'];
};


export type MutationCreateCommentArgs = {
  body: Scalars['String'];
  postID: Scalars['String'];
};


export type MutationCreatePostArgs = {
  input: CreatePostInput;
};


export type MutationCreateUserArgs = {
  input: UserCreds;
};


export type MutationFollowUserArgs = {
  username: Scalars['String'];
};


export type MutationLoginArgs = {
  input: UserCreds;
};


export type MutationSignedUploadUrlArgs = {
  input: SignedUploadInput;
};


export type MutationUnFollowUserArgs = {
  username: Scalars['String'];
};

export type Post = {
  __typename?: 'Post';
  body: Scalars['String'];
  camera: Scalars['String'];
  commentCount: Scalars['Int'];
  film: Scalars['String'];
  id: Scalars['String'];
  images: Array<Image>;
  lens: Scalars['String'];
  publishedAt?: Maybe<Scalars['Time']>;
  user: User;
};

export type Query = {
  __typename?: 'Query';
  getComments: Array<Comment>;
  getNewPosts: Array<Post>;
  getPost: Post;
  getPosts: Array<Post>;
  getUser: User;
  getUserPosts: Array<Post>;
  me?: Maybe<User>;
};


export type QueryGetCommentsArgs = {
  before?: InputMaybe<Scalars['Time']>;
  postID: Scalars['String'];
};


export type QueryGetNewPostsArgs = {
  before?: InputMaybe<Scalars['Time']>;
};


export type QueryGetPostArgs = {
  id: Scalars['String'];
};


export type QueryGetPostsArgs = {
  before?: InputMaybe<Scalars['Time']>;
};


export type QueryGetUserArgs = {
  username: Scalars['String'];
};


export type QueryGetUserPostsArgs = {
  before?: InputMaybe<Scalars['Time']>;
  username: Scalars['String'];
};

export type SignedUploadInput = {
  contentType: Scalars['String'];
  fileName: Scalars['String'];
};

export type User = {
  __typename?: 'User';
  following: Array<Follow>;
  id: Scalars['String'];
  username: Scalars['String'];
};

export type UserCreds = {
  password: Scalars['String'];
  username: Scalars['String'];
};

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
};

export type Mutation = {
  __typename?: 'Mutation';
  createUser: User;
  login: User;
};


export type MutationCreateUserArgs = {
  input: UserCreds;
};


export type MutationLoginArgs = {
  input: UserCreds;
};

export type Post = {
  __typename?: 'Post';
  body: Scalars['String'];
  user: User;
};

export type Query = {
  __typename?: 'Query';
  posts: Array<Post>;
};

export type User = {
  __typename?: 'User';
  username: Scalars['String'];
};

export type UserCreds = {
  password: Scalars['String'];
  username: Scalars['String'];
};

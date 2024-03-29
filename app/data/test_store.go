package data

import "time"

type TestStore struct {
	TestCreateUser             func(*User) error
	CreateUserCallCount        int
	TestCreateFollow           func(*Follow) error
	CreateFollowCallCount      int
	TestDeleteFollow           func(*Follow) error
	DeleteFollowCallCount      int
	TestGetUser                func(string) (*User, error)
	GetUserCallCount           int
	TestCreateAccessToken      func(*User) (*AccessToken, error)
	CreateAccessTokenCallCount int
	TestDeleteAccessToken      func(string) error
	DeleteAccessTokenCallCount int
	TestGetAccessToken         func(string) (*AccessToken, error)
	GetAccessTokenCallCount    int
	TestCreatePost             func(*Post) error
	CreatePostCallCount        int
	TestGetPosts               func(*time.Time) ([]*Post, error)
	GetPostsCallCount          int
	TestGetPost                func(string) (*Post, error)
	GetPostCallCount           int
	TestGetUserPosts           func(string, *time.Time) ([]*Post, error)
	GetUserPostsCallCount      int
	TestGetUsersPosts          func([]string, *time.Time) ([]*Post, error)
	GetUsersPostsCallCount     int
	TestGetUserFollows         func(string) ([]*Follow, error)
	GetUserFollowsCallCount    int
	TestCreateComment          func(*Comment) error
	CreateCommentCallCount     int
	TestGetComments            func(string, *time.Time) error
	GetCommentsCallCount       int
}

func (t *TestStore) CreateUser(user *User) error {
	t.CreateUserCallCount += 1
	return t.TestCreateUser(user)
}

func (t *TestStore) CreateFollow(follow *Follow) error {
	t.CreateFollowCallCount += 1
	return t.TestCreateFollow(follow)
}

func (t *TestStore) DeleteFollow(follow *Follow) error {
	t.DeleteFollowCallCount += 1
	return t.TestDeleteFollow(follow)
}

func (t *TestStore) GetUser(username string) (*User, error) {
	t.GetUserCallCount += 1
	return t.TestGetUser(username)
}

func (t *TestStore) CreateAccessToken(user *User) (*AccessToken, error) {
	t.CreateAccessTokenCallCount += 1
	return t.TestCreateAccessToken(user)
}

func (t *TestStore) DeleteAccessToken(id string) error {
	t.DeleteAccessTokenCallCount += 1
	return t.TestDeleteAccessToken(id)
}

func (t *TestStore) GetAccessToken(id string) (*AccessToken, error) {
	t.GetAccessTokenCallCount += 1
	return t.TestGetAccessToken(id)
}

func (t *TestStore) CreatePost(post *Post) error {
	t.CreatePostCallCount += 1
	return t.TestCreatePost(post)
}

func (t *TestStore) GetPosts(before *time.Time) ([]*Post, error) {
	t.GetPostsCallCount += 1
	return t.TestGetPosts(before)
}

func (t *TestStore) GetPost(id string) (*Post, error) {
	t.GetPostCallCount += 1
	return t.TestGetPost(id)
}

func (t *TestStore) GetOrCreateImage(string, int, int) (*Image, error) {
	panic("should not be called")
}

func (t *TestStore) PublishPost(string, []*Image) error {
	panic("should not be called")
}

func (t *TestStore) UnpublishPost(string) error {
	panic("should not be called")
}

func (t *TestStore) GetUserPosts(id string, before *time.Time) ([]*Post, error) {
	t.GetUserPostsCallCount += 1
	return t.TestGetUserPosts(id, before)
}

func (t *TestStore) GetUsersPosts(ids []string, before *time.Time) ([]*Post, error) {
	t.GetUsersPostsCallCount += 1
	return t.TestGetUsersPosts(ids, before)
}

func (t *TestStore) GetUserFollows(id string) ([]*Follow, error) {
	t.GetUserFollowsCallCount += 1
	return t.TestGetUserFollows(id)
}

func (t *TestStore) CreateComment(comment *Comment) error {
	t.CreateCommentCallCount += 1
	return t.TestCreateComment(comment)
}

func (t *TestStore) GetComments(postID string, before *time.Time) ([]*Comment, error) {
	t.GetCommentsCallCount += 1
	return t.GetComments(postID, before)
}

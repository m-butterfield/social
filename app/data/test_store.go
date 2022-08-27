package data

type TestStore struct {
	TestCreateUser             func(*User) error
	CreateUserCallCount        int
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
	TestGetPosts               func() ([]*Post, error)
	GetPostsCallCount          int
	TestGetPost                func(int) (*Post, error)
	GetPostCallCount           int
	TestGetUserPosts           func(int) ([]*Post, error)
	GetUserPostsCallCount      int
}

func (t *TestStore) CreateUser(user *User) error {
	t.CreateUserCallCount += 1
	return t.TestCreateUser(user)
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

func (t *TestStore) GetPosts() ([]*Post, error) {
	t.GetPostsCallCount += 1
	return t.TestGetPosts()
}

func (t *TestStore) GetPost(id int) (*Post, error) {
	t.GetPostCallCount += 1
	return t.TestGetPost(id)
}

func (t *TestStore) GetOrCreateImage(string, int, int) (*Image, error) {
	panic("should not be called")
}

func (t *TestStore) PublishPost(int, []*Image) error {
	panic("should not be called")
}

func (t *TestStore) GetUserPosts(id int) ([]*Post, error) {
	t.GetUserPostsCallCount += 1
	return t.TestGetUserPosts(id)
}
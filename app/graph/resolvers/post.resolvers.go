package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"errors"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/graph/model"
	"github.com/m-butterfield/social/app/lib"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostInput) (*data.Post, error) {
	user, err := loggedInUser(ctx)
	if err != nil {
		return nil, internalError(err)
	}
	if user == nil {
		return nil, unauthorizedError()
	}

	if len(input.Images) == 0 {
		return nil, errors.New("no images provided")
	}
	if len(input.Body) > 4096 {
		return nil, errors.New("post body too long (max 4096 characters)")
	}

	post := &data.Post{
		Body:   input.Body,
		UserID: user.ID,
	}
	if err = r.DS.CreatePost(post); err != nil {
		return nil, internalError(err)
	}

	if _, err := r.TC.CreateTask("publish_post", "social-publish-post", &lib.PublishPostRequest{
		PostID: post.ID,
		Images: input.Images,
	}); err != nil {
		return nil, internalError(err)
	}

	return post, nil
}

// GetPost is the resolver for the getPost field.
func (r *queryResolver) GetPost(ctx context.Context, id string) (*data.Post, error) {
	user, err := loggedInUser(ctx)
	if err != nil {
		return nil, internalError(err)
	}
	if user == nil {
		return nil, unauthorizedError()
	}

	post, err := r.DS.GetPost(id)
	if err != nil {
		return nil, internalError(err)
	}
	if post == nil {
		return nil, errors.New("post not found")
	}
	return post, nil
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context) ([]*data.Post, error) {
	user, err := loggedInUser(ctx)
	if err != nil {
		return nil, internalError(err)
	}
	if user == nil {
		posts, err := r.DS.GetPosts()
		if err != nil {
			return nil, internalError(err)
		}
		return posts, nil
	}
	follows, err := r.DS.GetUserFollows(user.ID)
	if err != nil {
		return nil, internalError(err)
	}
	userIDs := []string{user.ID}
	for _, follow := range follows {
		userIDs = append(userIDs, follow.FollowerID)
	}
	posts, err := r.DS.GetUsersPosts(userIDs)
	if err != nil {
		return nil, internalError(err)
	}
	return posts, nil
}

// GetNewPosts is the resolver for the getNewPosts field.
func (r *queryResolver) GetNewPosts(ctx context.Context) ([]*data.Post, error) {
	posts, err := r.DS.GetPosts()
	if err != nil {
		return nil, internalError(err)
	}
	return posts, nil
}

// GetUserPosts is the resolver for the getUserPosts field.
func (r *queryResolver) GetUserPosts(ctx context.Context, userID string) ([]*data.Post, error) {
	user, err := r.DS.GetUser(userID)
	if err != nil {
		return nil, internalError(err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	posts, err := r.DS.GetUserPosts(user.ID)
	if err != nil {
		return nil, internalError(err)
	}
	return posts, nil
}

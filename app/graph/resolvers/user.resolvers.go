package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"errors"
	"strings"

	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/graph/generated"
	"github.com/m-butterfield/social/app/graph/model"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserCreds) (*data.User, error) {
	input.Username = strings.TrimSpace(input.Username)
	if input.Username == "" {
		return nil, errors.New("please provide a username")
	}
	if len(input.Username) > 64 {
		return nil, errors.New("username must be less than 64 characters long")
	}
	if len(input.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters long")
	}
	if len(input.Password) > 64 {
		return nil, errors.New("password must be less than 64 characters long")
	}

	if exising, err := r.DS.GetUser(input.Username); err != nil {
		return nil, internalError(err)
	} else if exising != nil {
		return nil, errors.New("username already exists")
	}

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return nil, internalError(err)
	}
	user := &data.User{
		Username: input.Username,
		Password: string(hashedPW),
	}
	if err = r.DS.CreateUser(user); err != nil {
		return nil, internalError(err)
	}
	if err = cookieLogin(ctx, r.DS, user); err != nil {
		return nil, internalError(err)
	}
	return user, nil
}

// FollowUser is the resolver for the followUser field.
func (r *mutationResolver) FollowUser(ctx context.Context, username string) (bool, error) {
	follower, err := loggedInUser(ctx)
	if err != nil {
		return false, internalError(err)
	}
	if follower == nil {
		return false, errors.New("not logged in")
	}

	user, err := r.DS.GetUser(username)
	if err != nil {
		return false, internalError(err)
	}
	if user == nil {
		return false, errors.New("user not found")
	}
	for _, follow := range follower.Following {
		if follow.UserID == user.ID {
			return true, nil
		}
	}

	if err = r.DS.CreateFollow(&data.Follow{
		FollowerID: follower.ID,
		UserID:     user.ID,
	}); err != nil && !data.IsDuplicateKeyError(err) {
		return false, internalError(err)
	}
	return true, nil
}

// UnFollowUser is the resolver for the unFollowUser field.
func (r *mutationResolver) UnFollowUser(ctx context.Context, username string) (bool, error) {
	follower, err := loggedInUser(ctx)
	if err != nil {
		return false, internalError(err)
	}
	if follower == nil {
		return false, errors.New("not logged in")
	}

	user, err := r.DS.GetUser(username)
	if err != nil {
		return false, internalError(err)
	}
	if user == nil {
		return false, errors.New("user not found")
	}
	followerFound := false
	for _, follow := range follower.Following {
		if follow.UserID == user.ID {
			followerFound = true
			break
		}
	}
	if !followerFound {
		return true, nil
	}

	if err = r.DS.DeleteFollow(&data.Follow{
		FollowerID: follower.ID,
		UserID:     user.ID,
	}); err != nil {
		return false, internalError(err)
	}
	return true, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.UserCreds) (*data.User, error) {
	input.Username = strings.TrimSpace(input.Username)
	user, err := r.DS.GetUser(input.Username)
	if err != nil {
		return nil, internalError(err)
	}
	if user == nil {
		return nil, errors.New("invalid username")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid password")
	}

	if err = cookieLogin(ctx, r.DS, user); err != nil {
		return nil, internalError(err)
	}
	return user, nil
}

// Logout is the resolver for the logout field.
func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	gctx, err := ginContextFromContext(ctx)
	if err != nil {
		return false, internalError(err)
	}
	cookie, err := getSessionCookie(gctx.Request)
	if err != nil {
		return false, internalError(err)
	}
	if cookie == nil {
		return true, nil
	}
	if err := r.DS.DeleteAccessToken(cookie.Value); err != nil {
		return false, internalError(err)
	}
	unsetSessionCookie(gctx.Writer)
	return true, nil
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*data.User, error) {
	user, err := loggedInUser(ctx)
	if err != nil {
		return nil, internalError(err)
	}
	return user, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, username string) (*data.User, error) {
	user, err := r.DS.GetUser(username)
	if err != nil {
		return nil, internalError(err)
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

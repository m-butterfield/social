package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/graph/generated"
	"github.com/m-butterfield/social/app/graph/model"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*data.User, error) {
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

	hashedPW, err := bcrypt.GenerateFromPassword([]byte(input.Password), 8)
	if err != nil {
		return nil, err
	}
	user := &data.User{
		Username: input.Username,
		Password: string(hashedPW),
	}
	if err = r.DS.CreateUser(user); err != nil {
		return nil, err
	}
	if err = cookieLogin(ctx, r.DS, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*data.Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

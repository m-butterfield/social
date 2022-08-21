package controllers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/graph/generated"
	"github.com/m-butterfield/social/app/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*data.User, error) {
	user := data.User{Username: input.Username}
	if err := ds.CreateUser(&user); err != nil {
		return nil, err
	}
	return &user, nil
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

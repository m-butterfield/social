package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"errors"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/m-butterfield/social/app/graph/model"
	"github.com/m-butterfield/social/app/lib"
	"golang.org/x/oauth2/google"
)

// SignedUploadURL is the resolver for the signedUploadURL field.
func (r *mutationResolver) SignedUploadURL(ctx context.Context, input model.SignedUploadInput) (string, error) {
	if user, err := loggedInUser(ctx); err != nil {
		return "", internalError(err)
	} else if user == nil {
		return "", unauthorizedError()
	}

	if input.FileName == "" {
		return "", errors.New("no FileName provided")
	}
	if input.ContentType == "" {
		return "", errors.New("missing ContentType value")
	}
	fileName := lib.UploadsPrefix + input.FileName

	serviceAccount, err := os.ReadFile("/secrets/uploadercreds.json")
	if err != nil {
		return "", internalError(err)
	}
	conf, err := google.JWTConfigFromJSON(serviceAccount)
	if err != nil {
		return "", internalError(err)
	}
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type: " + input.ContentType,
		},
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().UTC().Add(15 * time.Minute),
	}

	url, err := storage.SignedURL(lib.ContentBucket, fileName, opts)
	if err != nil {
		return "", internalError(err)
	}
	return url, nil
}

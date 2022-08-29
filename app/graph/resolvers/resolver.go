package resolvers

import (
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DS data.Store
	TC lib.TaskCreator
}

package graph

import "github.com/m-butterfield/social/app/data"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DS data.Store
}

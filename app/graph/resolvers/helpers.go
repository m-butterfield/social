package resolvers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"net/http"
)

func ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(lib.GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

func cookieLogin(ctx context.Context, ds data.Store, user *data.User) error {
	token, err := ds.CreateAccessToken(user)
	if err != nil {
		return err
	}

	gctx, err := ginContextFromContext(ctx)
	if err != nil {
		return err
	}

	http.SetCookie(gctx.Writer, &http.Cookie{
		Name:    lib.SessionTokenName,
		Value:   token.ID,
		Expires: token.ExpiresAt,
	})
	return nil
}

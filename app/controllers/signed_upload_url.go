package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/lib"
	"golang.org/x/oauth2/google"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type signedUploadURLRequest struct {
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
}

type signedUploadURLResponse struct {
	URL string `json:"url"`
}

func signedUploadURL(c *gin.Context) {
	req := &signedUploadURLRequest{}
	err := c.Bind(req)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if req.FileName == "" {
		c.String(http.StatusBadRequest, "Please provide a file name")
		return
	}
	if req.ContentType == "" {
		c.String(http.StatusBadRequest, "Please provide the content type")
		return
	}
	fileName := lib.UploadsPrefix + req.FileName

	serviceAccount, err := os.ReadFile("/secrets/uploadercreds.json")
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	conf, err := google.JWTConfigFromJSON(serviceAccount)
	if err != nil {
		return
	}
	opts := &storage.SignedURLOptions{
		Scheme: storage.SigningSchemeV4,
		Method: "PUT",
		Headers: []string{
			"Content-Type: " + req.ContentType,
		},
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().UTC().Add(15 * time.Minute),
	}

	url, err := storage.SignedURL(lib.ContentBucket, fileName, opts)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	c.JSON(http.StatusOK, &signedUploadURLResponse{URL: url})
}

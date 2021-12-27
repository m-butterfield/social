package tasks

import (
	"cloud.google.com/go/storage"
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"image"
	_ "image/jpeg"
	"io"
	"log"
)

func publishPost(c *gin.Context) {
	createReq := &lib.PublishPostRequest{}
	err := c.Bind(createReq)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	var images []*data.Image
	for _, fileName := range createReq.Images {
		img, err := saveImage(fileName)
		if err != nil {
			lib.InternalError(err, c)
			return
		}
		images = append(images, img)
	}
	if err = ds.PublishPost(createReq.PostID, images); err != nil {
		lib.InternalError(err, c)
		return
	}
}

func saveImage(fileName string) (*data.Image, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer func(client *storage.Client) {
		if err := client.Close(); err != nil {
			log.Println(err)
		}
	}(client)

	bucket := client.Bucket(lib.ContentBucket)
	upload := bucket.Object(lib.UploadsPrefix + fileName)

	width, height, err := getDimensions(ctx, upload)
	if err != nil {
		return nil, err
	}

	hash, err := getHash(ctx, upload)
	if err != nil {
		return nil, err
	}

	result := bucket.Object(hash + ".jpg")
	if _, err := result.CopierFrom(upload).Run(ctx); err != nil {
		return nil, err
	}
	if err := result.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return nil, err
	}
	if err := upload.Delete(ctx); err != nil {
		return nil, err
	}

	return ds.GetOrCreateImage(result.ObjectName(), width, height)
}

func getDimensions(ctx context.Context, obj *storage.ObjectHandle) (int, int, error) {
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return 0, 0, err
	}
	defer func(reader *storage.Reader) {
		if err := reader.Close(); err != nil {
			log.Println(err)
		}
	}(reader)
	imgConf, _, err := image.DecodeConfig(reader)
	if err != nil {
		return 0, 0, err
	}
	return imgConf.Width, imgConf.Height, nil
}

func getHash(ctx context.Context, obj *storage.ObjectHandle) (string, error) {
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return "", err
	}
	defer func(reader *storage.Reader) {
		if err := reader.Close(); err != nil {
			log.Println(err)
		}
	}(reader)
	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

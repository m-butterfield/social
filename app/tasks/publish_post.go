package tasks

import (
	"cloud.google.com/go/storage"
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h2non/bimg"
	"github.com/m-butterfield/social/app/data"
	"github.com/m-butterfield/social/app/lib"
	"io/ioutil"
	"log"
	"math"
)

const (
	maxWidth  = 800
	maxHeight = 1000
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

	size, imgData, err := processImage(ctx, upload)
	if err != nil {
		return nil, err
	}

	hash, err := getHash(imgData)
	if err != nil {
		return nil, err
	}
	result := bucket.Object(hash + ".jpg")
	w := result.NewWriter(ctx)
	w.ContentType = "image/jpeg"
	if _, err = w.Write(imgData); err != nil {
		return nil, err
	}
	if err = w.Close(); err != nil {
		return nil, err
	}
	if err := result.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return nil, err
	}
	if err := upload.Delete(ctx); err != nil {
		return nil, err
	}

	return ds.GetOrCreateImage(result.ObjectName(), size.Width, size.Height)
}

func processImage(ctx context.Context, obj *storage.ObjectHandle) (*bimg.ImageSize, []byte, error) {
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer func(reader *storage.Reader) {
		if err := reader.Close(); err != nil {
			log.Println(err)
		}
	}(reader)
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, nil, err
	}

	img := bimg.NewImage(buffer)
	if _, err = img.AutoRotate(); err != nil {
		return nil, nil, err
	}

	size, err := img.Size()
	if err != nil {
		return nil, nil, err
	}
	width := size.Width
	height := size.Height

	if width > maxWidth {
		ratio := float64(height) / float64(width)
		width = maxWidth
		height = int(math.Round(float64(width) * ratio))
	}
	if height > maxHeight {
		ratio := float64(width) / float64(height)
		height = maxHeight
		width = int(math.Round(float64(height) * ratio))
	}

	imgData, err := img.Process(bimg.Options{
		Width:   width,
		Height:  height,
		Quality: 85,
	})
	if err != nil {
		return nil, nil, err
	}
	return &bimg.ImageSize{
		Width:  width,
		Height: height,
	}, imgData, nil
}

func getHash(data []byte) (string, error) {
	hash := sha256.New()
	if _, err := hash.Write(data); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

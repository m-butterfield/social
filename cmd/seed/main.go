package main

import (
	"github.com/m-butterfield/social/app/data"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var ds data.Store

func main() {
	if err := seedData(); err != nil {
		log.Fatal(err)
	}
}

func seedData() error {
	var err error
	if ds, err = data.Connect(); err != nil {
		return err
	}
	matt, err := addUser("matt")
	if err != nil {
		return err
	}
	caitlin, err := addUser("caitlin")
	if err != nil {
		return err
	}
	if err = addPost(
		caitlin.ID,
		"3bbed2c9ba9f7a9e7236a5ece4f8b3f133a21773f31caf9b27bdb5317f9c3145.jpg",
		750,
		1000,
		"Hank & Grump",
	); err != nil {
		return err
	}
	if err = addPost(
		caitlin.ID,
		"dcfb468c82a246d6162dc9ea041964e0463fcb053ec603f92964a4cb459c23e9.jpg",
		750,
		1000,
		"Littleneck",
	); err != nil {
		return err
	}
	if err = addPost(
		matt.ID,
		"52c692472d738924279fc47e6913607d9d1d2e6dc85413c2d5d00aa493438d9b.jpg",
		640,
		640,
		"another post",
	); err != nil {
		return err
	}
	if err = addPost(
		matt.ID,
		"ae67015c1136c8bbdaac9a6efe3f2f9f217bef6a6f48e6d782d1a7766a2b6c27.jpg",
		800,
		800,
		"nyc",
	); err != nil {
		return err
	}
	if err = addPost(
		matt.ID,
		"b7b055ac4d85954817d523c28bf1994efda563d87491b1058f7baa60418bed8d.jpg",
		750,
		1000,
		"city cat",
	); err != nil {
		return err
	}
	return nil
}

func addPost(userID, imageID string, width, height int, postBody string) error {
	post := &data.Post{
		UserID: userID,
		Body:   postBody,
	}
	if err := ds.CreatePost(post); err != nil {
		return err
	}
	image, err := ds.GetOrCreateImage(imageID, width, height)
	if err != nil {
		return err
	}
	return ds.PublishPost(post.ID, []*data.Image{image})
}

func addUser(userName string) (*data.User, error) {
	hashedPW, err := bcrypt.GenerateFromPassword([]byte("password"), 8)
	if err != nil {
		return nil, err
	}
	user := &data.User{
		ID:       userName,
		Password: string(hashedPW),
	}
	err = ds.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

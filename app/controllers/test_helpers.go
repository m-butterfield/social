package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/social/app/data"
)

func testRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := router()
	return r
}

type testStore struct {
	createUser func(*data.User) error
	getUser    func(string) (*data.User, error)
}

func (t *testStore) CreateUser(user *data.User) error {
	return t.createUser(user)
}

func (t *testStore) GetUser(id string) (*data.User, error) {
	return t.getUser(id)
}

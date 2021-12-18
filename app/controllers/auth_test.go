package controllers

import (
	"github.com/m-butterfield/social/app/data"
	"testing"
)

func TestAuthGoodToken(t *testing.T) {
	ts := &testStore{
		getAccessToken: func(id string) (*data.AccessToken, error) {
			return nil, nil
		},
	}
	ds = ts
}

func TestAuthBadToken(t *testing.T) {

}

func TestAuthNoToken(t *testing.T) {

}

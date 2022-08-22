package controllers

import (
	"github.com/m-butterfield/social/app/data"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLogout(t *testing.T) {
	w := httptest.NewRecorder()
	tokenID := "1234"
	ts := &data.TestStore{
		TestGetAccessToken: func(s string) (*data.AccessToken, error) {
			return &data.AccessToken{ID: tokenID}, nil
		},
		TestDeleteAccessToken: func(id string) error {
			assert.Equal(t, id, tokenID)
			return nil
		},
	}
	ds = ts
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.AddCookie(&http.Cookie{
		Name:  sessionTokenName,
		Value: tokenID,
	})

	testRouter().ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
	assert.Equal(t, ts.DeleteAccessTokenCallCount, 1)
	cookies := w.Result().Cookies()
	assert.Equal(t, len(cookies), 1)
	sessionCookie := cookies[0]
	assert.Equal(t, sessionCookie.Value, "")
	assert.Equal(t, sessionCookie.Expires, time.Unix(0, 0).UTC())
}

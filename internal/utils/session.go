package store

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func InitializeStore(secretKey []byte) {
	store = sessions.NewCookieStore(secretKey)
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return store.Get(r, "sessionid")
}

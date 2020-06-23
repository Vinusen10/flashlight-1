package middleware

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var store *sessions.CookieStore

//TODO Authentication handler
func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/login", http.StatusFound)
		} else {
			h(w, r)
		}
	}
}

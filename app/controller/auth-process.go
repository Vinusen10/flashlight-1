package controller

import (
	"crypto/rand"
	"flashlight/app/model"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var store *sessions.CookieStore

func init() {
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key := make([]byte, 32)
	rand.Read(key)
	store = sessions.NewCookieStore(key)
}
func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-login.html", "template/components/login.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "layout", nil)
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	check := model.CheckPassword(username, password)

	if check == true {
		session, _ := store.Get(r, "session")
		// Set user as authenticated
		session.Values["authenticated"] = true
		session.Values["username"] = username
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		t, err := template.ParseFiles("template/index-login.html", "template/components/login.html")
		if err != nil {
			log.Println(err)
		}
		t.ExecuteTemplate(w, "index-login.html", nil)
	}
}
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

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

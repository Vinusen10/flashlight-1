package controller

import (
	"crypto/rand"
	"flashlight/app/model"
	"fmt"
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

	// checking password
	check := model.CheckPassword(username, password)
	fmt.Println(check)
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
			log.Fatal(err)
		}
		t.ExecuteTemplate(w, "index-login.html", nil)
	}
}

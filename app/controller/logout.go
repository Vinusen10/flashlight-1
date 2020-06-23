package controller

import (
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello")
	http.Redirect(w, r, "/", http.StatusFound)
}

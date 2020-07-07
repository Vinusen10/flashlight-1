package controller

import (
	"log"
	"net/http"
)

func Like(w http.ResponseWriter, r *http.Request) {

}

func SendComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.Form["comment"]
	log.Print(comment)

	if r.Method == "GET" {
		//Execute Template
	}
	if r.Method == "POST" {
		//Parsing & Processing Form-Message
	}
	http.Redirect(w, r, "/logged", http.StatusFound)
}

package controller

import (
	"html/template"
	"log"
	"net/http"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-login.html", "template/components/login.html")
	if err != nil {
		log.Println(err)
	}
	//MockData
	t.ExecuteTemplate(w, "layout", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	//
}

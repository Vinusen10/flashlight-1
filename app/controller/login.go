package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		t, err := template.ParseFiles("template/index-login.html", "template/components/login.html")
		if err != nil {
			log.Println(err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	}
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println("Username:" + username + "| Password:" + password)
	} else {
		// print a error message or error page
	}
}

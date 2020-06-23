package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Logged(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-logged.html", "template/components/logged.html", "template/components/box-type-1.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "layout", nil)
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

package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Logged(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-logged.html", "template/components/logged.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "layout", nil)
}

package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/components/login.html")
	log.Println(err)
	t.ExecuteTemplate(w, "layout", nil)
	//MockData
}

package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-register.html", "template/components/register.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "layout", nil)
	//MockData
}

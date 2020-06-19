package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-upload.html", "template/components/upload.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "layout", nil)

}

package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-gallrey.html", "template/components/gallery.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "layout", nil)

}

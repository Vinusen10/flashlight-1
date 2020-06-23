package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/index-gallrey.html", "template/components/gallery.html", "template/components/card.html")
		if err != nil {
			log.Println(err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	} else {

	}
}

func DeletePicture(w http.ResponseWriter, r *http.Request) {

}

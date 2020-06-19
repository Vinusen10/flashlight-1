package controller

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/components/box.html", "template/components/landingHeader.html")
	if err != nil {
		log.Println(err)
	}

	//Mockdata

	t.ExecuteTemplate(w, "layout", nil)

}

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
	type Box struct {
	}

	t.ExecuteTemplate(w, "layout", Box{})

}

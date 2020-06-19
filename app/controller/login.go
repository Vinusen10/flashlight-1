package controller

import (
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
		//MockData
		t.ExecuteTemplate(w, "layout", nil)
	}

}

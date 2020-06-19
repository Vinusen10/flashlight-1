package controller

import (
	"github.com/gorilla/schema"
	"html/template"
	"log"
	"net/http"
)

var decoder = schema.NewDecoder()

type User struct {
	Name     string
	Password string
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/index-register.html", "template/components/register.html")
		if err != nil {
			log.Println(err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	} else {

		//Validation Process
		var user User

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}

		err = decoder.Decode(&user, r.PostForm)
		if err != nil {
			log.Println(err)
		}
	}

}

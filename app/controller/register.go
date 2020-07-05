package controller

import (
	"flashlight/app/model"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Name     string
	Password string
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("template/index-register.html", "template/components/register.html")
	if err != nil {
		log.Println(err)
	}
	t.ExecuteTemplate(w, "layout", nil)

}
func RegisterProcess(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	if model.UserExist(username) != true {
		t, err := template.ParseFiles("template/index-register.html", "template/components/register.html")
		if err != nil {
			log.Println(err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	}

}

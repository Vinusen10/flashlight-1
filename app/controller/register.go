package controller

import (
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
func Register(w http.ResponseWriter, r *http.Request) {

}

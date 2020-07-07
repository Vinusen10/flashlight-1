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

	usermail := r.FormValue("username")
	password := r.FormValue("password")

	if model.UserExist(usermail) != true {
		t, err := template.ParseFiles("template/index-register.html", "template/components/register.html")
		if err != nil {
			log.Println(err)
		}
		t.ExecuteTemplate(w, "layout", nil)
	} else {
		username := model.MailToUsername(usermail)
		u := model.User{}
		u.Password = password
		u.Mail = usermail
		u.Username = username

		err := u.AddUser()
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

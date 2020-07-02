package controller

import (
	"flashlight/app/model"
	"html/template"
	"log"
	"net/http"
)

func Logged(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-logged.html", "template/components/logged.html", "template/components/box.html")
	if err != nil {
		log.Println(err)
	}

	allPosts, _ := model.GetAllPosts()

	data := struct {
		Posts []map[string]interface{}
	}{
		Posts: allPosts,
	}
	t.ExecuteTemplate(w, "layout", data)
}

func SendComment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	comment := r.Form["comment"]
	log.Print(comment)

	if r.Method == "GET" {
		//Execute Template
	}
	if r.Method == "POST" {
		//Parsing & Processing Form-Message
	}
	http.Redirect(w, r, "/logged", http.StatusFound)
}

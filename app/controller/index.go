package controller

import (
	"flashlight/app/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		t, err := template.ParseFiles("template/index.html", "template/components/box.html", "template/components/landingHeader.html")
		if err != nil {
			log.Println(err)
		}

		allPosts, _ := model.GetAllPosts()

		indexData := struct {
			Posts []map[string]interface{}
		}{
			Posts: allPosts,
		}
		t.ExecuteTemplate(w, "layout", indexData)
	} else {
		t, err := template.ParseFiles("template/index-logged.html", "template/components/logged.html", "template/components/auth-user-box.html")
		if err != nil {
			log.Println(err)
		}
		u := session.Values["username"].(string)
		user, _ := model.GetUserByMail(u)
		allPosts, _ := model.GetAllPosts()

		loggedData := struct {
			Posts []map[string]interface{}
			User  model.User
		}{
			Posts: allPosts,
			User:  user,
		}
		t.ExecuteTemplate(w, "layout", loggedData)
	}
}

func SendComment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	user := session.Values["username"].(string)
	comment := r.FormValue("comment")
	id := r.FormValue("data")
	fmt.Println(id)
	fmt.Println(user)
	fmt.Println(comment)
	model.AppendComment(id, comment, user)

	http.Redirect(w, r, "/#"+id, http.StatusFound)
}

func Like(w http.ResponseWriter, r *http.Request) {
	//TODO
}

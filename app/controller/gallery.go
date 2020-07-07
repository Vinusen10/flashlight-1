package controller

import (
	"flashlight/app/model"
	"html/template"
	"log"
	"net/http"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index-gallrey.html", "template/components/gallery.html", "template/components/card.html")
	if err != nil {
		log.Println(err)
	}
	session, _ := store.Get(r, "session")
	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {

		user := session.Values["username"].(string)
		data, _ := model.GetPostsbyUser(user)

		gallery := struct {
			MyPosts []map[string]interface{}
		}{
			MyPosts: data,
		}
		t.ExecuteTemplate(w, "layout", gallery)
	}

}

func DeletePicture(w http.ResponseWriter, r *http.Request) {

}

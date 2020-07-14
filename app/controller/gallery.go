package controller

import (
	"flashlight/app/model"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, r, "/", http.StatusNotFound)

	} else {
		t, err := template.ParseFiles("template/index-gallrey.html", "template/components/gallery.html", "template/components/card.html")
		if err != nil {
			log.Println(err)
		}
		username := session.Values["username"].(string)
		user, _ := model.GetUserByMail(username)
		data, _ := model.GetPostsbyUser(username)
		gallery := struct {
			MyPosts []map[string]interface{}
			User    model.User
		}{
			MyPosts: data,
			User:    user,
		}
		t.ExecuteTemplate(w, "layout", gallery)
	}

}

func DeletePicture(w http.ResponseWriter, r *http.Request) {
	_id := r.FormValue("data")
	fmt.Println(_id)
	post, err := model.GetPostByID(_id)
	if err != nil {
		log.Println(err)
	}
	post.Delete()
	http.Redirect(w, r, "/gallery", http.StatusFound)
}

package controller

import (
	"flashlight/app/model"
	"html/template"
	"log"
	"net/http"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("template/index-gallrey.html", "template/components/gallery.html", "template/components/card.html")
		if err != nil {
			log.Println(err)
		}
		data, _ := model.GetPostsbyUser("Justus Jonas")
		gallery := struct {
			MyPosts []map[string]interface{}
		}{
			MyPosts: data,
		}
		t.ExecuteTemplate(w, "layout", gallery)
	} else {

	}
}

func DeletePicture(w http.ResponseWriter, r *http.Request) {

}

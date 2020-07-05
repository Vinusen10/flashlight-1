package controller

import (
	"flashlight/app/model"
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/components/box.html", "template/components/landingHeader.html")
	if err != nil {
		log.Println(err)
	}

	allPosts, _ := model.GetAllPosts()
	likes := 666
	indexData := struct {
		Posts []map[string]interface{}
		Likes int
	}{
		Posts: allPosts,
		Likes: likes,
	}

	t.ExecuteTemplate(w, "layout", indexData)
}

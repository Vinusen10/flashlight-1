package controller

import (
	"flashlight/app/model"
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

		allPosts, _ := model.GetAllPosts()

		loggedData := struct {
			Posts []map[string]interface{}
		}{
			Posts: allPosts,
		}
		t.ExecuteTemplate(w, "layout", loggedData)
	}

}

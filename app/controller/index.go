package controller

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/index.html", "template/components/box.html", "template/components/landingHeader.html")
	if err != nil {
		log.Println(err)
	}

	now := time.Now()
	//Mockdata
	mockData := struct {
		Name     string
		Date     string
		Likes    int
		Caption  string
		Comments []map[string]interface{} //TODO
	}{
		Name:    "Ghirishaanth Ananthavadivel",
		Date:    now.Format("02 Jan 06 15:04"),
		Likes:   10,
		Caption: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		//Comments -> Model JSON
	}

	t.ExecuteTemplate(w, "layout", mockData)

}

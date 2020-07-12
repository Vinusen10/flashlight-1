package controller

import (
	"flashlight/app/model"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Upload(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")
	if session.Values["authenticated"] != nil {
		t, err := template.ParseFiles("template/index-upload.html", "template/components/upload.html")
		if err != nil {
			log.Println(err)
		}
		uploadData := struct {
			Username string
		}{
			Username: session.Values["username"].(string),
		}
		t.ExecuteTemplate(w, "layout", uploadData)
	} else {
		http.Redirect(w, r, "/", http.StatusNotFound)
	}

}

func Uploading(w http.ResponseWriter, r *http.Request) {
	//Parsing input
	session, _ := store.Get(r, "session")

	r.ParseMultipartForm(10 << 20)

	// Retrieve file from posted form-data
	caption := r.FormValue("caption")
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	fmt.Println("Uploading File:", handler.Filename)
	fmt.Println("Uploading File:", handler.Size)
	fmt.Println("Uploading File:", handler.Header)

	// Write temporary file on the server
	tempFile, err := ioutil.TempFile("./content/", "upload-*.jpeg")
	if err != nil {
		log.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	post := model.Post{}
	post.Caption = caption
	post.Author = session.Values["username"].(string)
	post.ImagPath = tempFile.Name()
	post.Timestamp.Date = time.Now().Format("2006-01-02")
	post.Timestamp.Time = time.Now().Format("15:04")
	fmt.Println(post.Timestamp)
	post.Likes = []string{}
	post.Comments = []model.Comment{}
	post.AddPost()
	tempFile.Write(fileBytes)
	http.Redirect(w, r, "/", http.StatusFound)
}

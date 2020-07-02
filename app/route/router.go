package route

import (
	"flashlight/app/controller"
	"github.com/gorilla/mux"
	"net/http"
)

//GetRouter is a router that routes requests to their controller
func GetRouter() *mux.Router {
	r := mux.NewRouter()

	//controllers
	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/sendComment", controller.SendComment).Methods("POST")
	r.HandleFunc("/like", nil).Methods("POST")
	r.HandleFunc("/login", controller.LoginPage).Methods("GET")
	r.HandleFunc("/login-process", controller.Login).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("POST")
	r.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	r.HandleFunc("/register-process", controller.Register).Methods("POST")
	r.HandleFunc("/gallery", controller.Gallery).Methods("GET")
	r.HandleFunc("/deletePicture", controller.DeletePicture).Methods("POST")
	r.HandleFunc("/logged", controller.Logged) // TODO
	r.HandleFunc("/uploading", controller.Uploading).Methods("POST")
	r.HandleFunc("/uploads", controller.Upload).Methods("GET")

	//statics
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	//See all files on server
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files", http.FileServer(http.Dir("."))))
	//Content-Files
	r.PathPrefix("/content/").Handler(http.StripPrefix("/content", http.FileServer(http.Dir("./content/"))))

	return r
}

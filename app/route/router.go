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
	r.HandleFunc("/logout", controller.Logout)
	r.HandleFunc("/like", controller.Like)
	r.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	r.HandleFunc("/register-process", controller.Register).Methods("POST")
	r.HandleFunc("/gallery", controller.Gallery).Methods("GET")
	r.HandleFunc("/deletePicture", controller.DeletePicture).Methods("POST")
	r.HandleFunc("/logged", controller.Logged)
	r.HandleFunc("/uploads", controller.Upload)

	//statics
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files", http.FileServer(http.Dir("."))))

	return r
}

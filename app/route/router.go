package route

import (
	"flashlight/app/controller"
	"github.com/gorilla/mux"
	"net/http"
)

//GetRouter is a router that routes requests to their controller
func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/gallery", controller.Gallery).Methods("GET")
	r.HandleFunc("/uploads", controller.Upload).Methods("GET")
	r.HandleFunc("/authenticate", controller.AuthenticateUser)
	r.HandleFunc("/login", controller.Login).Methods("GET")
	r.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	r.HandleFunc("/registerProcess", controller.RegisterProcess).Methods("POST")
	r.HandleFunc("/logout", controller.Logout).Methods("GET")

	// Controller for Interactions
	r.HandleFunc("/sendComment", controller.Auth(controller.SendComment)).Methods("POST")
	r.HandleFunc("/like", nil).Methods("POST").Methods("POST")
	r.HandleFunc("/deletePicture", controller.Auth(controller.DeletePicture))
	r.HandleFunc("/uploading", controller.Auth(controller.Uploading)).Methods("POST")

	/*File Server & CND */
	//statics
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))
	//See all files on server
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files", http.FileServer(http.Dir("."))))
	//Content-Files
	r.PathPrefix("/content/").Handler(http.StripPrefix("/content", http.FileServer(http.Dir("./content/"))))

	return r
}

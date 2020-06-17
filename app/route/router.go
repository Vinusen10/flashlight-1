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
	r.HandleFunc("/", controller.Index)
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/gallery", controller.Gallery)
	r.HandleFunc("/", controller.Logged)
	r.HandleFunc("/uploads", controller.Upload)

	//statics
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./static/"))))

	return r
}

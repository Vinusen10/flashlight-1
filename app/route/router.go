package route

import (
	"flashlight/app/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/index", controller.Index).Methods("GET")

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("./assets/static/css"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("./assets/static/ressource/img"))))
	return r
}

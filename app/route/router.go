package route

import (
	"flashlight/app/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	//controllers
	r.HandleFunc("", controller.Index)
	r.HandleFunc("/login", controller.Login)
	r.HandleFunc("/register", controller.Register)
	r.HandleFunc("/gallery", controller.Gallery)
	r.HandleFunc("/", controller.Logged)
	r.HandleFunc("/uploads", controller.Upload)
	r.HandleFunc()

	//statics
	r.PathPrefix("/static/").Handler(http.StripPrefix("static", http.FileServer(http.Dir("./assets/template/pages"))))
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("./assets/static/css"))))
	r.PathPrefix("/img/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("./assets/static/ressource/img"))))

	return r
}

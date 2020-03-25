package router

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix()
	return r
}

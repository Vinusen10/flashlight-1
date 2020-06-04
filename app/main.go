package main

import (
	"flashlight/app/route"
	"log"
	"net/http"
	"time"
)

func main() {

	server := &http.Server{
		Addr:           ":80",
		Handler:        route.GetRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}

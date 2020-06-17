package main

import (
	"flashlight/app/route"
	"log"
	"net/http"
	"time"
)

func main() {

	r := route.GetRouter()
	server := &http.Server{
		Handler:      r,
		Addr:         "localhost:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Server is Running...")
	log.Fatal(server.ListenAndServe())
}

package main

import (
	"flashlight/app/route"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	r := route.GetRouter()
	server := &http.Server{
		Handler:      r,
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Server is Running...")
	fmt.Printf("%-55s", fmt.Sprintf("%55s", "http://"+server.Addr))

	log.Fatal(server.ListenAndServe())
}

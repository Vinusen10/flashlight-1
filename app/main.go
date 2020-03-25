package main

import (
	"log"
	"net/http"
)

func main(){


	server := &http.Server{
		Addr:"localhost:8000",
	}
	log.Fatal(server.ListenAndServe())

}

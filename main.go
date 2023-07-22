package main

/*
 * Author: Cyprien Borée
 * Email: cyprien[dot]boree[at]tuta[dot]io
 * */

import (
	"fmt"
	"log"
	"net/http"

	"url-aliaser/handler"
	"url-aliaser/model"
)

func main() {

	// endpoints
	http.HandleFunc("/alias", handler.HandleShortenRequest)
	http.HandleFunc("/", handler.HandleRedirectionRequest)

	// start the server on port 8080
	fmt.Printf("Server listening on port %s:\n", model.ServerPort)
	log.Fatal(http.ListenAndServe(":"+model.ServerPort, nil))
}

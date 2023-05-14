package main

import (
	"fmt"
	"log"
	"net/http"

	"url-shortener/handler"
)

func main() {

	http.HandleFunc("/shorten", handler.HandleRequest)

	// Start the server on port 8080
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

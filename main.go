package main

import (
	"log"
	"net/http"
)

var NewUrl string = "https://github.com"

func main() {
	mux := http.NewServeMux()

	rh := http.RedirectHandler(NewUrl, http.StatusTemporaryRedirect)

	mux.Handle("/", rh)

	log.Print("Listening...")

	http.ListenAndServe(":3000", mux)
}

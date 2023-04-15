package main

import (
	"log"
	"net/http"

	"github.com/boreec/URL-shortener/handler"
)

func main() {
	mux := http.NewServeMux()

	urlShortenerHandler := handler.URLShortenerHandler{}
	homepageHandler := handler.HomepageHandler{}

	mux.Handle("/", &homepageHandler)
	mux.Handle("/shorten", &urlShortenerHandler)

	log.Print("Listening...")

	http.ListenAndServe(":3000", mux)
}

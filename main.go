package main

import (
	"log"
	"net/http"

	"github.com/boreec/URL-shortener/handler"
)

func main() {
	mux := http.NewServeMux()

	url_shortener := handler.URLShortenerHandler{}

	mux.Handle("/shorten", &url_shortener)

	log.Print("Listening...")

	http.ListenAndServe(":3000", mux)
}

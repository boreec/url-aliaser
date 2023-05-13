package main

import (
	"log"
	"net/http"

	"github.com/boreec/URL-shortener/handler"
)

func main() {
	mux := http.NewServeMux()

	urlShortenerHandler := handler.NewURLShortenerHandler()
	homepageHandler := handler.NewHomepageHandler()

	mux.Handle("/", homepageHandler)
	mux.Handle("/shorten", urlShortenerHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	log.Print("Listening on localhost:3000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf(err.Error())
	}
}

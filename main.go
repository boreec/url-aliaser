package main

import (
	"log"
	"net/http"
)

var NewUrl string = "https://github.com"

func main() {
	http.Handle("/", http.RedirectHandler(NewUrl, http.StatusMovedPermanently))
	log.Fatal(http.ListenAndServe(":9000", nil))
}

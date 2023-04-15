package handler

import "net/http"

type URLShortenerHandler struct{}

func (url_shortener *URLShortenerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world !"))
}

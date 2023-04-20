package handler

import (
	"fmt"
	"net/http"
)

type URLShortenerHandler struct {
	// urlMap is a map that matches a long url given by the user to a shortened url calculated by a hash function.
	urlMap map[string]string
}

// create a new URLShortenerHandler with an empty urlMap.
func NewURLShortenerHandler() *URLShortenerHandler {
	return &URLShortenerHandler{urlMap: make(map[string]string)}
}

func (url_shortener *URLShortenerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	url := r.Form.Get("long_link")

	// check if url is already contained in urlList
	if url_shortener.urlMap[url] != "" {
		w.Write([]byte(fmt.Sprintf("link %s is already shortened as %s", url, url_shortener.urlMap[url])))
	} else {
		url_shortener.urlMap[url] = url_shortener.hash(url)
	}

	w.Write([]byte(fmt.Sprintf("to do: shorten %v", url_shortener)))
}

// create a hash of a given string
func (url_shortener *URLShortenerHandler) hash(url string) string {
	// to do
	return "/shortenedlink"
}

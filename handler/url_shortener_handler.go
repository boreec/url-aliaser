package handler

import (
	"fmt"
	"net/http"
)

type URLShortenerHandler struct {
	urlMap map[string]string
}

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

func (url_shortener *URLShortenerHandler) hash(url string) string {
	// to do
	return "/shortenedlink"
}

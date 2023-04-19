package handler

import (
	"fmt"
	"net/http"
)

type URLShortenerHandler struct {
	urlList map[string]string
}

func (url_shortener *URLShortenerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	url := r.Form.Get("long_link")

	// check if url is already contained in urlList
	if url_shortener.urlList[url] == "" {
		w.Write([]byte(fmt.Sprintf("link %s is already shortened as %s", url, url_shortener.urlList[url])))
	} else {
		url_shortener.urlList[url] = shorten(url)
	}

	w.Write([]byte(fmt.Sprintf("to do: shorten %v", url_shortener)))
}

package handler

import (
	"fmt"
	"net/http"
)

type URLShortenerHandler struct {
	longURL string
}

func (url_shortener *URLShortenerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	url_shortener.longURL = r.Form.Get("long_link")
	w.Write([]byte(fmt.Sprintf("to do: shorten %v", url_shortener)))
}

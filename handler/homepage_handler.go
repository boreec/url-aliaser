package handler

import "net/http"

type HomepageHandler struct{}

func (hp *HomepageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
	// to do: display the list of handlers
}

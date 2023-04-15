package handler

import "net/http"

type HomepageHandler struct{}

func (hp *HomepageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// to do: display the list of handlers
}

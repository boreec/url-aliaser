package handler

/*
 * Author: Cyprien Borée
 * Email: cyprien[dot]boree[at]tuta[dot]io
 * */

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"url-aliaser/model"
)

var ErrInvalidRequestMethod = errors.New("invalid request method")
var urlMap = make(map[string]string) // Map to store the shortened URLs

// expected payload in the request
type PayloadRequest struct {
	Url    string `json:"url"`              // required
	Length uint16 `json:"length,omitempty"` // optional
}

// payload returned
type PayloadResponse struct {
	Url string `json:"url"`
}

func HandleShortenRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, ErrInvalidRequestMethod.Error(), http.StatusMethodNotAllowed)
		return
	}

	// decode request's payload into a PayloadRequest
	var payloadRequest PayloadRequest
	err := json.NewDecoder(r.Body).Decode(&payloadRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// shorten url
	hash, err := model.ShortenURL(payloadRequest.Url, payloadRequest.Length)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	aliasUrl := model.ServerPath + "/" + hash
	urlMap[aliasUrl] = payloadRequest.Url

	log.Printf("bind url '%s' to '%s'", payloadRequest.Url, aliasUrl)

	sendResponse(w, PayloadResponse{Url: aliasUrl})
}

func sendResponse(w http.ResponseWriter, payload PayloadResponse) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	// response body
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleRedirectionRequest(w http.ResponseWriter, r *http.Request) {
	requestUrl := model.ServerPath + r.URL.Path
	originalUrl, exists := urlMap[requestUrl]
	if !exists {
		log.Printf("request url '%s' not found", requestUrl)
		http.NotFound(w, r)
		return
	}
	log.Printf("Redirecting '%s' to '%s'", requestUrl, originalUrl)
	http.Redirect(w, r, originalUrl, http.StatusFound)
}

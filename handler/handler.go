package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"url-shortener/model"
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

func HandleRequest(w http.ResponseWriter, r *http.Request) {
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

	// Check if the URL is already shortened
	if shortUrl, ok := urlMap[payloadRequest.Url]; ok {
		// URL already exists in the map, return the existing shortened URL
		sendResponse(w, PayloadResponse{Url: shortUrl})
		return
	}

	// shorten url
	shortUrl, err := model.ShortenURL(payloadRequest.Url, payloadRequest.Length)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	urlMap[payloadRequest.Url] = shortUrl

	sendResponse(w, PayloadResponse{Url: shortUrl})
}

func sendResponse(w http.ResponseWriter, payload PayloadResponse) {
	// response header
	w.Header().Set("Content-Type", "application/json")

	// response body
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

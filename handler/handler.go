package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"url-shortener/model"
)

var ErrInvalidRequestMethod = errors.New("invalid request method")

// expected payload in the request
type PayloadRequest struct {
	Url    string `json:"url"`              // required
	Length string `json:"length,omitempty"` // optional
}

// payload returned
type PayloadResponse struct {
	Url string `json:"url"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, ErrInvalidRequestMethod.Error(), http.StatusMethodNotAllowed)
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
	shortUrl, err := model.ShortenUrl(payloadRequest.Length)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// return response with payload within
	var payloadResponse PayloadResponse
	payloadResponse.Url = shortUrl
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payloadResponse)
}

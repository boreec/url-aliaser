package handler

import (
	"encoding/json"
	"net/http"
)

// expected payload in the request
type PayloadRequest struct {
	Url    string `json:"url"`              // required
	Length string `json:"length,omitempty"` // optional
}

// payload returned
type PayloadResponse struct {
	Url string `json:"url"`
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var payloadRequest PayloadRequest
		err := json.NewDecoder(r.Body).Decode(&payloadRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Return the JSON payload as the response
		var payloadResponse PayloadResponse
		payloadResponse.Url = "to do"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(payloadResponse)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

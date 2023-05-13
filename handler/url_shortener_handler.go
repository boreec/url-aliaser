package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"net/http"

	"golang.org/x/exp/slog"
)

var (
	ErrEmptyGivenUrl = errors.New("given url is empty")
)

type URLShortenerHandler struct {
	// urlMap is a map that matches a long url given by the user to a shortened url calculated by a hash function.
	urlMap map[string]string
}

// create a new URLShortenerHandler with an empty urlMap.
func NewURLShortenerHandler() *URLShortenerHandler {
	return &URLShortenerHandler{urlMap: make(map[string]string)}
}

func WriteError(w http.ResponseWriter, err error) {
	slog.Info(err.Error())
	w.Write([]byte(err.Error()))
}

type RequestPayload struct {
	Url       string `json:"url"`                  // required
	MaxLength int    `json:"max_length,omitempty"` // empty
}

type ResponsePayload struct {
	Url string `json:"url"`
}

func (url_shortener *URLShortenerHandler) Handler(w http.ResponseWriter, r *http.Request) {

	// parse incoming json payload
	var request RequestPayload
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		// Handle parsing error
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// to do
	// Construct the response payload
	response := ResponsePayload{
		Url: "shortened url",
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the response payload
	if err := json.NewEncoder(w).Encode(response); err != nil {
		// Handle encoding error
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

// create a hash of a given string
func (url_shortener *URLShortenerHandler) hash(url string, length int) (string, error) {
	if url == "" {
		return "", ErrEmptyGivenUrl
	}

	hash := fnv.New32a()
	hash.Write([]byte(url))
	hashValue := hash.Sum32()

	// convert the hash value to hexadecimal string
	hashString := fmt.Sprintf("%0*x", length/2, hashValue)

	return hashString[:length], nil
}

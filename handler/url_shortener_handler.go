package handler

import (
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
)

type URLShortenerHandler struct {
	// urlMap is a map that matches a long url given by the user to a shortened url calculated by a hash function.
	urlMap map[string]string
}

// create a new URLShortenerHandler with an empty urlMap.
func NewURLShortenerHandler() *URLShortenerHandler {
	return &URLShortenerHandler{urlMap: make(map[string]string)}
}

func (url_shortener *URLShortenerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	err = r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	url := r.Form.Get("long_link")

	// check if url is already contained in urlList
	if url_shortener.urlMap[url] != "" {
		w.Write([]byte(fmt.Sprintf("link %s is already shortened as %s", url, url_shortener.urlMap[url])))
	} else {
		if url_shortener.urlMap[url], err = url_shortener.hash(url, 10); err != nil {
			log.Printf("failed to shorten link '%s' for the following error:\n%s", url, err.Error())
			w.Write([]byte(fmt.Sprintf("failed to shorten link '%s' for the following error:\n%s", url, err.Error())))
		}
	}

	w.Write([]byte(fmt.Sprintf("to do: shorten %v", url_shortener)))
}

// create a hash of a given string
func (url_shortener *URLShortenerHandler) hash(url string, length int) (string, error) {
	if url == "" {
		return "", errors.New("url to hash can not be empty")
	}

	hash := fnv.New32a()
	hash.Write([]byte(url))
	hashValue := hash.Sum32()

	// convert the hash value to hexadecimal string
	hashString := fmt.Sprintf("%0*x", length/2, hashValue)

	return hashString[:length], nil
}

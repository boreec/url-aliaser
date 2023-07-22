package handler

/*
 * Author: Cyprien Borée
 * Email: cyprien[dot]boree[at]tuta[dot]io
 * */

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleRequestFailsOnGETRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a ResponseRecorder to record the handler response
	rr := httptest.NewRecorder()

	// Call the handler function with the mock request and response recorder
	HandleShortenRequest(rr, req)

	// Check status
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Result().StatusCode)

	// Check content of error
	assert.Equal(t, ErrInvalidRequestMethod.Error(), strings.TrimRight(rr.Body.String(), "\n"))
}

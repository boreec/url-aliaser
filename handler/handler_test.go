package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleRequestFailsOnGETRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a ResponseRecorder to record the handler response
	rr := httptest.NewRecorder()

	// Call the handler function with the mock request and response recorder
	HandleRequest(rr, req)

	// Check if the handler returned an error
	assert.Equal(t, http.StatusMethodNotAllowed, rr.Result().StatusCode)
}

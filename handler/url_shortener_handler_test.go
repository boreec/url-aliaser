package handler

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashWithEmptyString(t *testing.T) {
	shortenerHandler := NewURLShortenerHandler()

	_, err := shortenerHandler.hash("", 10)

	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, ErrEmptyGivenUrl))
}

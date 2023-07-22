package model

/*
 * Author: Cyprien Borée
 * Email: cyprien[dot]boree[at]tuta[dot]io
 * */

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenUrlCorrectLength(t *testing.T) {
	url := "https://example.com"

	for i := 1; i < 9; i++ {
		shortened_url, err := ShortenURL(url, uint16(i))
		assert.NoError(t, err)
		assert.Equal(t, i, len(shortened_url))
	}
}

func TestShortenUrlLengthZero(t *testing.T) {
	url := "https://example.com"
	_, err := ShortenURL(url, 0)
	assert.ErrorIs(t, err, ErrURLLengthZero)
}

func TestShortenUrlTooLong(t *testing.T) {
	url := "https://example.com"
	_, err := ShortenURL(url, URLMaxShorteningLength+1)
	assert.ErrorIs(t, err, ErrURLLengthTooLong)
}

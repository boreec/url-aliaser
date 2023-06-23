package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenUrlCorrectLength(t *testing.T) {
	url := "https://example.com"

	for i := 1; i < 9; i++ {
		shortened_url, err := ShortenUrl(url, 1)
		assert.NoError(t, err)
		assert.Equal(t, len(shortened_url), i)
	}
}

func TestShortenUrlLengthZero(t *testing.T) {
	url := "https://example.com"
	_, err := ShortenUrl(url, 0)
	assert.ErrorIs(t, err, ErrUrlLengthZero)
}

func TestShortenUrlTooLong(t *testing.T) {
	url := "https://example.com"
	_, err := ShortenUrl(url, UrlMaxLength+1)
	assert.ErrorIs(t, err, ErrUrlLengthTooLong)
}

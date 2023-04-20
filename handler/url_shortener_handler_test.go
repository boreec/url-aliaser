package handler

import "testing"

func TestHashWithEmptyString(t *testing.T) {
	shortenerHandler := NewURLShortenerHandler()

	if _, err := shortenerHandler.hash("", 10); err == nil {
		t.Fatal("hash function should return error when empty string\n")
	}
}

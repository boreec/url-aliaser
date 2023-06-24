package model

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"
)

var (
	ErrURLLengthZero     = errors.New("provided Length for shortening can not 0")
	ErrURLLengthTooLong  = errors.New("provided Length for shortening can not be this long")
	ErrURLWrongFormat    = errors.New("provided URL for shortening has incorrect format")
	ErrURLNotHttpOrHttps = errors.New("provided URL for shortening is not http or https")
)

const (
	URLMaxLength = 2048
)

// for any given url, provide another one shorter
func ShortenUrl(rawURL string, length uint16) (string, error) {

	if err := validateLength(length); err != nil {
		return "", err
	}

	if err := validateURL(rawURL); err != nil {
		return "", err
	}

	hashedURL := hash(rawURL, length)

	return hashedURL, nil
}

func hash(rawURL string, length uint16) string {
	hasher := sha256.New()
	hasher.Write([]byte(rawURL))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)[:length]
}

// validateURL checks if a given string represents a well-formed URL.
//
// Returns `nil` if the URL has a correct format and scheme.
// Returns `ErrURLWrongFormat` if the URL has an incorrect format.
// Returns `ErrURLMissingScheme` if the URL is missing the required http or https scheme.
func validateURL(rawUrl string) error {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return ErrURLWrongFormat
	}
	if !(parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
		return ErrURLNotHttpOrHttps
	}
	return nil
}

// validateLength checks if a given length can be used as a target to shorten a string.
//
// Returns `nil` if the length is an appropriate value.
// Returns `ErrURLLengthZero` if the given length is 0.
// Returns `ErrURLLengthTooLoog` if the given length exceeds URLMaxLength.
func validateLength(length uint16) error {
	if length == 0 {
		return ErrURLLengthZero
	}

	if length > URLMaxLength {
		return ErrURLLengthTooLong
	}
	return nil
}

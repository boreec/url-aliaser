package model

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"
)

var (
	ErrUrlLengthZero     = errors.New("provided Length for shortened url can not 0")
	ErrUrlLengthTooLong  = errors.New("provided Length for shortened url can not be this long")
	ErrUrlWrongFormat    = errors.New("provided Url for shortened url has incorrect format")
	ErrUrlNotHttpOrHttps = errors.New("provided Url is not http or https")
)

const (
	UrlMaxLength = 2048
)

// for any given url, provide another one shorter
func ShortenUrl(rawUrl string, length uint16) (string, error) {

	if err := validateLength(length); err != nil {
		return "", err
	}

	if err := validateURL(rawUrl); err != nil {
		return "", err
	}

	hashedUrl := hash(rawUrl, length)

	return hashedUrl, nil
}

func hash(rawUrl string, length uint16) string {
	hasher := sha256.New()
	hasher.Write([]byte(rawUrl))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)[:length]
}

// ValidateURL checks if a given string represents a well-formed URL.
//
// Returns `nil` if the URL has a correct format and scheme.
// Returns `ErrURLWrongFormat` if the URL has an incorrect format.
// Returns `ErrURLMissingScheme` if the URL is missing the required http or https scheme.
func validateURL(rawUrl string) error {
	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return ErrUrlWrongFormat
	}
	if !(parsedURL.Scheme == "http" || parsedURL.Scheme == "https") {
		return ErrUrlNotHttpOrHttps
	}
	return nil
}

func validateLength(length uint16) error {
	if length == 0 {
		return ErrUrlLengthZero
	}

	if length > UrlMaxLength {
		return ErrUrlLengthTooLong
	}
	return nil
}

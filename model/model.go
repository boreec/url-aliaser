package model

import (
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
func ShortenUrl(rawUrl string, length int) (string, error) {
	if length == 0 {
		return "", ErrUrlLengthZero
	}

	if length > UrlMaxLength {
		return "", ErrUrlLengthTooLong
	}

	// Validate the url
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", ErrUrlWrongFormat
	}
	if !(parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https") {
		return "", ErrUrlNotHttpOrHttps
	}

	return "to do", nil
}

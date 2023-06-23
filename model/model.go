package model

import "errors"

var (
	ErrUrlLengthZero    = errors.New("provided Length for shortened url can not 0")
	ErrUrlLengthTooLong = errors.New("provided Length for shortened url can not be this long")
)

const (
	UrlMaxLength = 2048
)

// for any given url, provide another one shorter
func ShortenUrl(url string, length int) (string, error) {
	return "to do", errors.New("to do")
}

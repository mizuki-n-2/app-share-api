package model

import (
	"errors"
	"regexp"
)

// Value Object: アプリのURL
type PostAppURL string

func NewPostAppURL(value string) (PostAppURL, error) {
	URL_PATTERN := `https?://[\w!\?/\+\-_~=;\.,\*&@#\$%\(\)'\[\]]+`
	
	if len(value) > 0 && !regexp.MustCompile(URL_PATTERN).MatchString(value) {
		return "", errors.New("urlの形式が正しくありません")
	}

	return PostAppURL(value), nil
}
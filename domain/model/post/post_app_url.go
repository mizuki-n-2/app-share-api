package model

import (
	"errors"
	"regexp"
)

// Value Object: アプリのURL
type PostAppURL string

func NewPostAppURL(value string) (PostAppURL, error) {
	urlPattern := `https?://[\w!\?/\+\-_~=;\.,\*&@#\$%\(\)'\[\]]+`
	if len(value) > 0 && !regexp.MustCompile(urlPattern).MatchString(value) {
		return "", errors.New("urlの形式が正しくありません")
	}

	return PostAppURL(value), nil
}
package model

import (
	"errors"
	"unicode/utf8"
)

// Value Object: 投稿内容
type PostContent string

func NewPostContent(value string) (PostContent, error) {
	if utf8.RuneCountInString(value) > 255 {
		return "", errors.New("contentは255文字以下にしてください")
	}

	return PostContent(value), nil
}
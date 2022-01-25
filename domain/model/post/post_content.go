package model

import (
	"fmt"
	"unicode/utf8"
)

// Value Object: 投稿内容
type PostContent string

func NewPostContent(value string) (PostContent, error) {
	MAX_LENGTH_POST_CONTENT := 255

	if utf8.RuneCountInString(value) > MAX_LENGTH_POST_CONTENT {
		return "", fmt.Errorf("contentは%d文字以下にしてください", MAX_LENGTH_POST_CONTENT)
	}

	return PostContent(value), nil
}
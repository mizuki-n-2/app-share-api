package model

import (
	"fmt"
	"unicode/utf8"
)

// Value Object: 投稿タイトル
type PostTitle string

func NewPostTitle(value string) (PostTitle, error) {
	MIN_LENGTH_POST_TITLE := 1
	MAX_LENGTH_POST_TITLE := 100

	if utf8.RuneCountInString(value) < MIN_LENGTH_POST_TITLE || utf8.RuneCountInString(value) > MAX_LENGTH_POST_TITLE {
		return "", fmt.Errorf("titleは%d文字以上%d文字以下にしてください", MIN_LENGTH_POST_TITLE, MAX_LENGTH_POST_TITLE)
	}

	return PostTitle(value), nil
}
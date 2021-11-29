package model

import (
	"errors"
	"unicode/utf8"
)

// Value Object: 投稿タイトル
type PostTitle string

func NewPostTitle(value string) (PostTitle, error) {
	if utf8.RuneCountInString(value) < 0 || utf8.RuneCountInString(value) > 100 {
		return "", errors.New("titleは1文字以上100文字以下にしてください")
	}

	return PostTitle(value), nil
}
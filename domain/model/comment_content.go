package model

import (
	"errors"
	"unicode/utf8"
)

// Value Object: コメント内容
type CommentContent string

func NewCommentContent(value string) (CommentContent, error) {
	if utf8.RuneCountInString(value) > 255 {
		return "", errors.New("contentは255文字以下にしてください")
	}

	return CommentContent(value), nil
}
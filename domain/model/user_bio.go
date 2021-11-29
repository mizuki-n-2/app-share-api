package model

import (
	"errors"
	"unicode/utf8"
)

// Value Object: 自己紹介文
type UserBio string

func NewUserBio(value string) (UserBio, error) {
	if utf8.RuneCountInString(value) > 255 {
		return "", errors.New("bioは255文字以下にしてください")
	}

	return UserBio(value), nil
}
package model

import (
	"fmt"
	"unicode/utf8"
)

// Value Object: 自己紹介文
type UserBio string

func NewUserBio(value string) (UserBio, error) {
	MAX_LENGTH_USER_BIO := 255

	if utf8.RuneCountInString(value) > MAX_LENGTH_USER_BIO {
		return "", fmt.Errorf("bioは%d文字以下にしてください", MAX_LENGTH_USER_BIO)
	}

	return UserBio(value), nil
}
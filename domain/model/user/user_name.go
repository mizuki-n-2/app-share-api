package model

import (
	"fmt"
	"unicode/utf8"
)

// Value Object: ユーザー名
type UserName string

func NewUserName(value string) (UserName, error) {
	MIN_LENGTH_USER_BIO := 2
	MAX_LENGTH_USER_BIO := 20

	if utf8.RuneCountInString(value) < MIN_LENGTH_USER_BIO || utf8.RuneCountInString(value) > MAX_LENGTH_USER_BIO {
		return "", fmt.Errorf("nameは%d文字以上%d文字以下にしてください", MIN_LENGTH_USER_BIO, MAX_LENGTH_USER_BIO)
	}

	return UserName(value), nil
}
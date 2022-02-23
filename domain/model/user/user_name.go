package model

import (
	"fmt"
	"unicode/utf8"
)

// Value Object: ユーザー名
type UserName string

func NewUserName(value string) (UserName, error) {
	MIN_LENGTH_USER_NAME := 2
	MAX_LENGTH_USER_NAME := 20

	if utf8.RuneCountInString(value) < MIN_LENGTH_USER_NAME || utf8.RuneCountInString(value) > MAX_LENGTH_USER_NAME {
		return "", fmt.Errorf("nameは%d文字以上%d文字以下にしてください", MIN_LENGTH_USER_NAME, MAX_LENGTH_USER_NAME)
	}

	return UserName(value), nil
}
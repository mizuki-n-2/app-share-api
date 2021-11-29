package model

import (
	"errors"
	"unicode/utf8"
)

// Value Object: ユーザー名
type UserName string

func NewUserName(value string) (UserName, error) {
	if utf8.RuneCountInString(value) < 1 || utf8.RuneCountInString(value) > 20 {
		return "", errors.New("nameは2文字以上20文字以下にしてください")
	}

	return UserName(value), nil
}
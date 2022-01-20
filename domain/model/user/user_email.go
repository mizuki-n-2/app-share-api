package model

import (
	"errors"
	"regexp"
)

// Value Object: メールアドレス
type UserEmail string

func NewUserEmail(value string) (UserEmail, error) {
	EMAIL_PATTERN := `^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	
	if !regexp.MustCompile(EMAIL_PATTERN).MatchString(value) {
		return "", errors.New("emailの形式が正しくありません")
	}

	return UserEmail(value), nil
}
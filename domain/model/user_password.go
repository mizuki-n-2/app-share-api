package model

import (
	"errors"
	"unicode/utf8"
	"golang.org/x/crypto/bcrypt"
)

// Value Object: パスワード
type UserPassword string

func NewUserPassword(value string) (UserPassword, error) {
	if utf8.RuneCountInString(value) < 8 || utf8.RuneCountInString(value) > 30 {
		return "", errors.New("passwordは8文字以上30文字以下にしてください")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}		

	return UserPassword(hashedPassword), nil
}
package model

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

// Value Object: パスワード
type UserPassword string

func NewUserPassword(value string) (UserPassword, error) {
	MIN_LENGTH_USER_PASSWORD := 8
	MAX_LENGTH_USER_PASSWORD := 30

	if utf8.RuneCountInString(value) < MIN_LENGTH_USER_PASSWORD || utf8.RuneCountInString(value) > MAX_LENGTH_USER_PASSWORD {
		return "", fmt.Errorf("passwordは%d文字以上%d文字以下にしてください", MIN_LENGTH_USER_PASSWORD, MAX_LENGTH_USER_PASSWORD)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}		

	return UserPassword(hashedPassword), nil
}
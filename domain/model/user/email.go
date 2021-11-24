package user

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(s string) (*Email, error) {
	emailPattern := `^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`
	if !regexp.MustCompile(emailPattern).MatchString(s) {
		return nil, errors.New("emailの形式が正しくありません")
	}

	email := Email(s)

	return &email, nil
}
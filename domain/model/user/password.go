package user

import (
	"errors"
	"unicode/utf8"
	"golang.org/x/crypto/bcrypt"
)

type Password string

func NewPassword(s string) (*Password, error) {
	if utf8.RuneCountInString(s) < 8 {
		return nil, errors.New("passwordは8文字以上にしてください")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}		
	
	password := Password(hashedPassword)

	return &password, nil
}

func (p *Password) Compare(s string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*p), []byte(s))
	if err != nil {
		return errors.New("パスワードが違います")
	}
	
	return nil
}
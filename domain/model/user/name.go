package user

import (
	"errors"
	"unicode/utf8"
)

type Name string

func NewName(s string) (*Name, error) {
	if utf8.RuneCountInString(s) <= 0 || utf8.RuneCountInString(s) > 10 {
		return nil, errors.New("nameは1文字以上10文字以下にしてください")
	}

	name := Name(s)

	return &name, nil
}
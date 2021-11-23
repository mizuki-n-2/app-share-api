package model

import (
	"time"
	"errors"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// userのコンストラクタ
func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("nameを入力してください")
	}
	if email == "" {
		return nil, errors.New("emailを入力してください")
	}
	if password == "" {
		return nil, errors.New("passwordを入力してください")
	}

	user := &User{
		Name:      name,
		Email:     email,
		Password:  password,
	}

	return user, nil
}

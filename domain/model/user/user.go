package user

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      Name      `json:"name"`
	Email     string    `json:"email"`
	Password  Password  `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	newName, err := NewName(name)
	if err != nil {
		return nil, err
	}

	if email == "" {
		return nil, errors.New("emailを入力してください")
	}

	newPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     *newName,
		Email:    email,
		Password: *newPassword,
		CreatedAt: time.Now(),
	}

	return user, nil
}

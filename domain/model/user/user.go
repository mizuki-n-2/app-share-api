package user

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      Name      `json:"name"`
	Email     Email     `json:"email"`
	Password  Password  `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	newName, err := NewName(name)
	if err != nil {
		return nil, err
	}

	newEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	newPassword, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Name:      *newName,
		Email:     *newEmail,
		Password:  *newPassword,
		CreatedAt: time.Now(),
	}

	return user, nil
}

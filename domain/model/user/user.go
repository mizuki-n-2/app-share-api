package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Entity: ユーザー
type User struct {
	ID        string       `json:"id"`
	Name      UserName     `json:"name" gorm:"not null;type:varchar(20)"`
	Email     UserEmail    `json:"email" gorm:"unique_index;not null"`
	Password  UserPassword `json:"password" gorm:"not null"`
	Avatar    string       `json:"avatar"`
	Bio       UserBio      `json:"bio" gorm:"size:255"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
}

func (user *User) ComparePassword(value string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(value))
	if err != nil {
		return errors.New("パスワードが違います")
	}

	return nil
}

func NewUser(name, email, password string) (*User, error) {
	userID := uuid.NewString()

	userName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	userEmail, err := NewUserEmail(email)
	if err != nil {
		return nil, err
	}

	userPassword, err := NewUserPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:     userID,
		Name:     userName,
		Email:    userEmail,
		Password: userPassword,
	}

	return user, nil
}

func (user *User) SetName(name string) error {
	userName, err := NewUserName(name)
	if err != nil {
		return err
	}

	user.Name = userName

	return nil
}

func (user *User) SetBio(bio string) error {
	userBio, err := NewUserBio(bio)
	if err != nil {
		return err
	}

	user.Bio = userBio

	return nil
}

func (user *User) SetAvatar(avatar string) error {
	// TODO: バリデーション
	user.Avatar = avatar

	return nil
}

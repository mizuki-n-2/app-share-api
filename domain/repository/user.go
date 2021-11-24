package repository

import (
	"app-share-api/domain/model/user"
)

type UserRepository interface {
	Store(user *user.User) (*user.User, error)
	FindByID(ID int) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}
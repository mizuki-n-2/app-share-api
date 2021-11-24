package repository

import (
	"app-share-api/domain/model"
)

type UserRepository interface {
	Store(user *model.User) (*model.User, error)
	FindByID(ID int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
package repository

import (
	"app-share-api/domain/model"
)

type UserRepository interface {
	Store(user *model.User) (*model.User, error)
	FindByID(id int) (*model.User, error)
}
package repository

import (
	"app-share-api/domain/model/user"
)

type UserRepository interface {
	Store(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	FindByID(ID string) (*model.User, error)
	FindByEmail(email model.UserEmail) (*model.User, error)
}
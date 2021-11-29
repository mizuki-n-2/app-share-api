package repository

import (
	"app-share-api/domain/model"
)

type LikeRepository interface {
	Store(like *model.Like) (*model.Like, error)
	FindByID(ID string) (*model.Like, error)
	Delete(ID string) error
}
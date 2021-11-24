package repository

import (
	"app-share-api/domain/model"
)

type LikeRepository interface {
	Store(like *model.Like) (*model.Like, error)
	FindByID(id int) (*model.Like, error)
	Delete(id int) error
}
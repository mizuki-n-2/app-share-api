package repository

import (
	"app-share-api/domain/model/like"
)

type LikeRepository interface {
	Store(like *like.Like) (*like.Like, error)
	FindByID(ID int) (*like.Like, error)
	Delete(ID int) error
}
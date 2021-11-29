package repository

import (
	"app-share-api/domain/model"
)

type PostRepository interface {
	Store(post *model.Post) (*model.Post, error)
	FindByID(ID string) (*model.Post, error)
	Update(post *model.Post) (*model.Post, error)
	Delete(ID string) error
}
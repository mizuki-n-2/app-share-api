package repository

import (
	"app-share-api/domain/model"
)

type PostRepository interface {
	Store(post *model.Post) (*model.Post, error)
	FindByID(ID int) (*model.Post, error)
	FindAll() ([]*model.Post, error)
	Delete(ID int) error
}
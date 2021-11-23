package repository

import (
	"app-share-api/domain/model"
)

type PostRepository interface {
	Store(post *model.Post) (*model.Post, error)
	FindByID(id int) (*model.Post, error)
	FindAll() ([]*model.Post, error)
	Delete(id int) error
}
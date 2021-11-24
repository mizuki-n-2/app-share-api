package repository

import (
	"app-share-api/domain/model/post"
)

type PostRepository interface {
	Store(post *post.Post) (*post.Post, error)
	FindByID(ID int) (*post.Post, error)
	FindAll() ([]*post.Post, error)
	Delete(ID int) error
}
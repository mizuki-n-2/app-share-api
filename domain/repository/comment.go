package repository

import (
	"app-share-api/domain/model"
)

type CommentRepository interface {
	Store(comment *model.Comment) (*model.Comment, error)
	FindByID(ID int) (*model.Comment, error)
	FindByPostID(postID int) ([]*model.Comment, error)
	Delete(ID int) error
}
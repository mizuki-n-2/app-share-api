package repository

import (
	"app-share-api/domain/model/comment"
)

type CommentRepository interface {
	Store(comment *model.Comment) (*model.Comment, error)
	FindByID(ID string) (*model.Comment, error)
	Update(comment *model.Comment) (*model.Comment, error)
	Delete(ID string) error
}

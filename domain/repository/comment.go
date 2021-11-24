package repository

import (
	"app-share-api/domain/model/comment"
)

type CommentRepository interface {
	Store(comment *comment.Comment) (*comment.Comment, error)
	FindByID(ID int) (*comment.Comment, error)
	FindByPostID(postID int) ([]*comment.Comment, error)
	Delete(ID int) error
}

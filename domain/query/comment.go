package query

import (
	"app-share-api/domain/model"
)

type CommentQueryService interface {
	FindByPostID(postID string) ([]model.Comment, error)
}

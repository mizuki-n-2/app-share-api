package query

import "app-share-api/domain/query/dto"
type CommentQueryService interface {
	GetCommentsByPostID(postID string) ([]*dto.Comment, error)
}

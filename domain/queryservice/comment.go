package queryservice

import "app-share-api/domain/queryservice/dto"
type CommentQueryService interface {
	GetCommentsByPostID(postID string) ([]*dto.Comment, error)
}

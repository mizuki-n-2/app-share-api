package queryservice

import "app-share-api/application/queryservice/dto"

type CommentQueryService interface {
	GetCommentsByPostID(postID string) ([]*dto.Comment, error)
}

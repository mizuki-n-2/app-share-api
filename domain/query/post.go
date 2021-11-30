package query

import "app-share-api/domain/query/dto"

type PostQueryService interface {
	GetAllPosts() ([]*dto.Post, error)
	GetPostByID(id string) (*dto.Post, error)
	GetPostsByUserID(userID string) ([]*dto.Post, error)
	GetLikePostsByUserID(userID string) ([]*dto.Post, error)
}
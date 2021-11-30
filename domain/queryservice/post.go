package queryservice

import "app-share-api/domain/queryservice/dto"

type PostQueryService interface {
	GetAllPosts() ([]*dto.Post, error)
	GetPostByID(id string) (*dto.Post, error)
	GetPostsByUserID(userID string) ([]*dto.Post, error)
	GetLikePostsByUserID(userID string) ([]*dto.Post, error)
}
package queryservice

import "app-share-api/domain/queryservice/dto"

type UserQueryService interface {
	GetUserByID(userID string) (*dto.User, error)
	GetRankingLikeUsers() ([]*dto.RankingLikeUser, error)
	GetRankingPostUsers() ([]*dto.RankingPostUser, error)
}
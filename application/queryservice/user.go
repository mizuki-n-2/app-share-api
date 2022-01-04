package queryservice

import "app-share-api/application/queryservice/dto"

type UserQueryService interface {
	GetUserByID(userID string) (*dto.User, error)
	GetRankingLikeUsers() ([]*dto.RankingLikeUser, error)
	GetRankingPostUsers() ([]*dto.RankingPostUser, error)
}
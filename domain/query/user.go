package query

import "app-share-api/domain/query/dto"

type UserQueryService interface {
	GetUserByID(userID string) (*dto.User, error)
	GetAllUsers() ([]*dto.User, error)
}
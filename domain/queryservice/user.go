package queryservice

import "app-share-api/domain/queryservice/dto"

type UserQueryService interface {
	GetUserByID(userID string) (*dto.User, error)
	GetAllUsers() ([]*dto.User, error)
}
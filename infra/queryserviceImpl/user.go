package queryserviceImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/queryservice"
)

type userQueryService struct {
	db *gorm.DB
}

func NewUserQueryService(db *gorm.DB) queryservice.UserQueryService {
	return &userQueryService{db: db}
}

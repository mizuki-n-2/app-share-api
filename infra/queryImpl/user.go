package queryImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/query"
)

type userQueryService struct {
	db *gorm.DB
}

func NewUserQueryService(db *gorm.DB) query.UserQueryService {
	return &userQueryService{db: db}
}

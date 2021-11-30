package queryserviceImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/queryservice"
)

type postQueryService struct {
	db *gorm.DB
}

func NewPostQueryService(db *gorm.DB) queryservice.PostQueryService {
	return &postQueryService{db: db}
}
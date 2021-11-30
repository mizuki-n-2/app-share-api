package queryserviceImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/queryservice"
)

type likeQueryService struct {
	db *gorm.DB
}

func NewLikeQueryService(db *gorm.DB) queryservice.LikeQueryService {
	return &likeQueryService{db: db}
}

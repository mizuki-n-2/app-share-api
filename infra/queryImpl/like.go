package queryImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/query"
)

type likeQueryService struct {
	db *gorm.DB
}

func NewLikeQueryService(db *gorm.DB) query.LikeQueryService {
	return &likeQueryService{db: db}
}

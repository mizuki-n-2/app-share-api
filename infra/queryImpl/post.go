package queryImpl

import (
	"gorm.io/gorm"

	"app-share-api/domain/query"
)

type postQueryService struct {
	db *gorm.DB
}

func NewPostQueryService(db *gorm.DB) query.PostQueryService {
	return &postQueryService{db: db}
}